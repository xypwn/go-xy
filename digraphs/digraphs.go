// Package digraphs provides utilities for directed graphs, represented as
// a mapping from node keys to edges.
package digraphs

import (
	"bytes"
	"fmt"
	"slices"
	"strings"

	"github.com/xypwn/go-xy/text"
)

func Reachable[K comparable](roots []K, edges func(K) []K) map[K]struct{} {
	reachable := map[K]struct{}{}
	nodes := slices.Clone(roots)
	var newNodes []K
	for len(nodes) > 0 {
		for _, node := range nodes {
			if _, ok := reachable[node]; ok {
				continue
			}
			reachable[node] = struct{}{}
			newNodes = append(newNodes, edges(node)...)
		}
		nodes, newNodes = newNodes, nodes[:0]
	}
	return reachable
}

// DOTCode generates graphviz DOT code to visualize a graph.
// nodes represents all nodes to be included in the visualization.
// name is the name of the digraph.
// prelude is DOT code inserted in the beginning.
// nodeAttrs should return a string representing a node's attributes
// (including the []).
func DOTCode[K comparable](nodes []K, edges func(K) []K, name, prelude string, nodeAttrs func(K) string) []byte {
	prelude = text.IndentString(strings.TrimSpace(prelude), "  ", 1)

	var b bytes.Buffer
	fmt.Fprintf(&b, "digraph %v {\n", name)
	if prelude != "" {
		b.WriteString(prelude)
		b.WriteByte('\n')
	}
	nodeIDs := map[K]int{}
	for id, key := range nodes {
		fmt.Fprintf(&b, "  %v", id)
		if attrs := nodeAttrs(key); attrs != "" {
			b.WriteByte(' ')
			b.WriteString(attrs)
		}
		b.WriteByte('\n')
		nodeIDs[key] = id
	}
	for id, key := range nodes {
		edgs := edges(key)
		edgs = slices.DeleteFunc(edgs, func(k K) bool {
			_, ok := nodeIDs[key]
			return !ok
		})
		if len(edgs) == 0 {
			continue
		}
		fmt.Fprintf(&b, "  %v -> {", id)
		for i, edg := range edgs {
			if i != 0 {
				b.WriteByte(' ')
			}
			fmt.Fprintf(&b, "%v", nodeIDs[edg])
		}
		fmt.Fprintf(&b, "}\n")
	}
	fmt.Fprintf(&b, "}\n")
	return b.Bytes()
}
