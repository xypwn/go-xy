package profile

import (
	"bytes"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

type Profile struct {
	lock     sync.Mutex
	buf      *bytes.Buffer // may or may not be used depending on the profiler
	outFile  string        // the file should only be created after Stop() is called
	close    func()
	timedOut bool
}

// CPU starts writing a CPU profile.
// Default output filename is "cpu.prof".
// Use [Profile.Stop] to stop profiling.
// Typical usage would be:
//
//	defer profile.CPU().Stop()
//
// Any errors will cause this function or
// [Profile.Stop] to panic.
func CPU() *Profile {
	p := &Profile{
		buf:     bytes.NewBuffer(nil),
		outFile: "cpu.prof",
	}
	if err := pprof.StartCPUProfile(p.buf); err != nil {
		panic("starting CPU profile: " + err.Error())
	}
	p.close = func() {
		pprof.StopCPUProfile()
		if err := os.WriteFile(p.outFile, p.buf.Bytes(), 0666); err != nil {
			panic("writing CPU profile: " + err.Error())
		}
	}
	return p
}

// Mem prepares memory profiling.
// Default output filename is "mem.prof".
// Use [Profile.Stop] to stop profiling.
//
// Unlike [CPU] profiling, this will consider the ENTIRE PROGRAM,
// no matter at which point of execution Mem was called.
//
// Any errors will cause this function or
// [Profile.Stop] to panic.
func Mem() *Profile {
	p := &Profile{
		outFile: "mem.prof",
	}
	p.close = func() {
		f, err := os.Create(p.outFile)
		if err != nil {
			panic("creating memory profile: " + err.Error())
		}
		defer func() {
			if err := f.Close(); err != nil {
				panic("closing memory profile file: " + err.Error())
			}
		}()
		runtime.GC()
		if err := pprof.Lookup("allocs").WriteTo(f, 0); err != nil {
			panic("writing memory profile: " + err.Error())
		}
	}
	return p
}

// Timeout will stop profiling after the specified duration,
// if [Profile.Stop] hasn't been called until then.
// Can be useful for debugging infinite loops if you don't know
// where they originate.
//
// The value pointed to by p is mutated - the input is returned as-is
// for convenience.
func (p *Profile) Timeout(duration time.Duration) *Profile {
	go func() {
		time.Sleep(duration)

		p.lock.Lock()
		defer p.lock.Unlock()

		p.timedOut = true
		p.close()
		p.close = func() {}
	}()
	return p
}

// Output sets the profile output filename.
//
// The value pointed to by p is mutated - the input is returned as-is
// for convenience.
func (p *Profile) Output(filename string) *Profile {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.outFile = filename
	return p
}

// Then will run fn after the profile has been stopped, either
// through timeout or [Stop].
//
// The value pointed to by p is mutated - the input is returned as-is
// for convenience.
func (p *Profile) Then(fn func(timedOut bool)) *Profile {
	close := p.close
	p.close = func() {
		close()
		fn(p.timedOut)
	}
	return p
}

// Stop will end profiling and flush file writes to the output file.
//
// It is valid to call this more than once, and from different goroutines.
func (p *Profile) Stop() {
	p.lock.Lock()
	defer p.lock.Unlock()

	p.close()
	p.close = func() {}
}
