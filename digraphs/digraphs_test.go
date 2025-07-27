package digraphs_test

import (
	"fmt"
	"maps"
	"slices"
	"strconv"

	"github.com/xypwn/go-xy/digraphs"
)

func ExampleReachable() {
	nodes := map[string][]string{
		"A": {"B", "C"},
		"B": {"D"},
		"C": nil,
		"D": {"A"},
		"E": {"A", "B"},
	}
	roots := []string{"A", "B"}
	reachable := digraphs.Reachable(
		roots,
		func(k string) []string { return nodes[k] },
	)
	fmt.Println(slices.Sorted(maps.Keys(reachable)))
	// Output: [A B C D]
}

func ExampleDOTCode() {
	nodes := map[string][]string{
		"A": {"B", "C"},
		"B": {"D"},
		"C": nil,
		"D": {"A"},
		"E": {"A", "B"},
	}
	includedNodes := []string{"A", "B", "D", "E"}
	code := digraphs.DOTCode(
		includedNodes,
		func(k string) []string { return nodes[k] },
		"my_graph",                      // name
		"node[shape=box, style=filled]", // prelude code
		func(k string) string { return fmt.Sprintf("[label=%v]", strconv.Quote(k)) }, // attributes
	)
	fmt.Println(string(code))
	// Output:
	// digraph my_graph {
	//   node[shape=box, style=filled]
	//   0 [label="A"]
	//   1 [label="B"]
	//   2 [label="D"]
	//   3 [label="E"]
	//   0 -> {1 0}
	//   1 -> {2}
	//   2 -> {0}
	//   3 -> {0 1}
	// }
}
