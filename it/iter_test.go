package it_test

import (
	"maps"
	"slices"
	"testing"

	"github.com/xypwn/go-xy/it"
	"github.com/xypwn/go-xy/tests"
)

func TestIter(tt *testing.T) {
	t := tests.New(tt)
	t.Run("join-map-sprint", func(t *tests.T) {
		s := slices.Values([]int{1, 2, 3, 4})
		t.Equal(
			"1, 2, 3, 4",
			it.Join(it.Map(s, it.Sprint), ", "),
		)
	})
	t.Run("uniq", func(t *tests.T) {
		s := slices.Values([]int{1, 1, 0, 2, 2, 3, 4, 4, 4})
		t.Equal(
			[]int{1, 0, 2, 3, 4},
			slices.Collect(it.Uniq(s)),
		)
	})
	t.Run("markovian", func(t *tests.T) {
		naturals := it.Markovian(1, func(x int) int { return x + 1 })
		t.Equal(
			[]int{1, 2, 3, 4, 5},
			slices.Collect(it.Take(naturals, 5)),
		)
	})
	t.Run("sorted-by-key", func(t *tests.T) {
		m := map[string]int{
			"zero":  0,
			"one":   1,
			"two":   2,
			"three": 3,
			"four":  4,
			"five":  5,
			"six":   6,
			"seven": 7,
			"eight": 8,
			"nine":  9,
		}
		t.Equal(
			[]int{8, 5, 4, 9, 1, 7, 6, 3, 2, 0},
			slices.Collect(it.Second(it.SortedByKey(m))),
		)
	})
	t.Run("merge", func(t *tests.T) {
		s1 := slices.Values([]int{1, 2, 3, 4})
		s2 := slices.Values([]int{4, 3, 2, 1})
		s := it.Merge(s1, s2)
		t.Equal(slices.Collect(s1), slices.Collect(it.First(s)))
		t.Equal(slices.Collect(s2), slices.Collect(it.Second(s)))
	})
	t.Run("repeat", func(t *tests.T) {
		t.Equal(
			[]int{1, 1, 1, 1},
			slices.Collect(it.Repeat(1, 4)),
		)
	})
	t.Run("with-first/second", func(t *tests.T) {
		firstFiveNaturals := it.Take(it.Markovian(1, func(x int) int { return x + 1 }), 5)
		t.Equal(
			map[int]string{1: "1", 2: "2", 3: "3", 4: "4", 5: "5"},
			maps.Collect(it.WithSecond(firstFiveNaturals, it.Sprint)),
		)
		t.Equal(
			map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5},
			maps.Collect(it.WithFirst(firstFiveNaturals, it.Sprint)),
		)
	})
}
