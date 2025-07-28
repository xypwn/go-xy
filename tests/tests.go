package tests

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/xypwn/go-xy/text"
)

type T struct {
	*testing.T
	DefaultPrinter func(any) string
}

func New(t *testing.T) *T {
	return &T{T: t}
}

func (t *T) Run(name string, run func(t *T)) {
	t.Helper()
	t.T.Run(name, func(tt *testing.T) {
		t.Helper()
		run(New(tt))
	})
}

func (t *T) Equal(expect, got any) {
	t.Helper()
	t.EqualPrinter(expect, got, t.DefaultPrinter)
}

func DefaultPrinter(value any) string {
	return fmt.Sprint(value)
}

// QuotePrinter quotes strings to make
// special characters and whitespaces
// legible.
// Also adds line numbers.
func QuotePrinter(value any) string {
	ilog10 := func(x int) int {
		v := -1
		for x > 0 {
			x /= 10
			v++
		}
		return v
	}

	s := fmt.Sprint(value)
	if strings.Contains(s, "\n") {
		var b strings.Builder
		sp := strings.Split(s, "\n")
		for i, line := range sp {
			lineNo := fmt.Sprintf("%-*v| ", ilog10(len(sp))+1, i+1)
			if i != 0 {
				b.WriteString("+\n")
			}
			if i != len(sp)-1 {
				line += "\n"
			}
			b.WriteString(lineNo + strconv.Quote(line))
		}
		return b.String()
	} else {
		return strconv.Quote(s)
	}
}

func (t *T) EqualPrinter(expect, got any, printer func(any) string) {
	t.Helper()
	if !reflect.DeepEqual(expect, got) {
		multiline := false
		expectStr := printer(expect)
		gotStr := printer(got)
		if strings.Contains(expectStr, "\n") {
			expectStr = text.IndentString(expectStr, "    ", 1)
			multiline = true
		}
		if strings.Contains(gotStr, "\n") {
			gotStr = text.IndentString(gotStr, "    ", 1)
			multiline = true
		}
		if multiline {
			t.Fatalf("Equal():\nexpected:\n%v\nbut got:\n%v\n", expectStr, gotStr)
		} else {
			t.Fatalf("Equal(): expected %v, but got %v", expectStr, gotStr)
		}
	}
}
