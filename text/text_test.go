package text

import (
	"testing"

	"github.com/xypwn/go-xy/tests"
)

func TestIndentString(tt *testing.T) {
	t := tests.New(tt)

	t.Equal(`  Hello
  World`,
		IndentString(`Hello
World`, "  ", 1),
	)

	t.Equal(`  Hello
  World
`,
		IndentString(`Hello
World
`, "  ", 1),
	)

	t.Equal(`  Hello
  World
  `,
		IndentString(`Hello
World
  `, "  ", 1),
	)

	t.Equal(`  Hello
  
  World
`,
		IndentString(`Hello
  
World
`, "  ", 1),
	)
}
