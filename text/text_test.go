package text_test

import (
	"fmt"
	"testing"

	"github.com/xypwn/go-xy/tests"
	"github.com/xypwn/go-xy/text"
)

func TestIndentString(tt *testing.T) {
	t := tests.New(tt)
	t.DefaultPrinter = tests.QuotePrinter

	t.Equal("  Hello\n  World",
		text.IndentString("Hello\nWorld", "  ", 1))

	t.Equal("  Hello\n  World\n",
		text.IndentString("Hello\nWorld\n", "  ", 1))

	t.Equal("  Hello\n  World\n",
		text.IndentString("Hello\nWorld\n  ", "  ", 1))

	t.Equal("  Hello\n\n  World\n",
		text.IndentString("Hello\n  \nWorld\n", "  ", 1))

	t.Equal(" A\r\n\r\n B\r\n",
		text.IndentString("A\r\n \r\nB\r\n", " ", 1))

func ExampleIndentString() {
	s := `
{
	x := "world"

	fmt.Println("Hello,", x)
}`
	fmt.Println(text.IndentString(s, "\t", 1))
	// Output:
	// 	{
	// 		x := "world"
	//
	// 		fmt.Println("Hello,", x)
	// 	}
}
