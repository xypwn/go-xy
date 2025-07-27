package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/xypwn/go-xy/profile"
)

// Usage:
//
//	go run ./examples/profileme && go tool pprof -http : cpu.prof
func main() {
	// For getting started:
	//defer profile.CPU().Stop()

	// This demonstrates a slightly more complex use case:
	defer profile.CPU().
		Timeout(5 * time.Second).
		Then(func(timedOut bool) {
			if timedOut {
				fmt.Println("Profiler timed out")
			}
			fmt.Println("Profile written to cpu.prof")
			fmt.Println("Open in browser with: go tool pprof -http : cpu.prof")
			if timedOut {
				os.Exit(0) // you usually don't want this - here it makes the demo easier to follow
			}
		}).
		Stop()

	// You can also run a memory profile (note that the memory profile
	// output is called mem.prof!):
	//defer profile.Mem().Stop()

	var jsonBlob = []byte(`[
	{"Name": "Platypus", "Order": "Monotremata"},
	{"Name": "Quoll",    "Order": "Dasyuromorphia"}
]`)
	type Animal struct {
		Name  string
		Order string
	}

	// Set this to true to see the timeout kick in.
	infinite := false

	for i := 0; infinite || i < 200_000; i++ {
		var a []Animal
		if err := json.Unmarshal(jsonBlob, &a); err != nil {
			log.Fatal(err)
		}
	}
}
