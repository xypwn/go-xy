package it

import (
	"cmp"
	"fmt"
	"iter"
	"maps"
	"slices"
	"strings"
)

// Map returns a new sequence where every element
// in seq is replaced by the result of putting
// that element through fn.
func Map[A, B any](seq iter.Seq[A], fn func(A) B) iter.Seq[B] {
	return func(yield func(B) bool) {
		for a := range seq {
			if !yield(fn(a)) {
				return
			}
		}
	}
}

// Fold accepts a start value b and applies fn to that value
// for every element in seq.
func Fold[A, B any](seq iter.Seq[A], b B, fn func(A, B) B) B {
	for a := range seq {
		b = fn(a, b)
	}
	return b
}

// SortedByKey returns a sequence that iterates over
// the given map with all items sorted by their key.
func SortedByKey[K cmp.Ordered, V any, M ~map[K]V](m M) iter.Seq2[K, V] {
	ks := slices.Sorted(maps.Keys(m))
	return func(yield func(K, V) bool) {
		for _, k := range ks {
			v := m[k]
			if !yield(k, v) {
				return
			}
		}
	}
}

// SortedByKeyFunc is like [SortedByKey], but accepts a comparison function.
// compare should return the same values [cmp.Compare].
// Note that the order will only be consistent if compare never compares equal
// for two keys in the map.
func SortedByKeyFunc[K comparable, V any, M ~map[K]V](m M, compare func(K, K) int) iter.Seq2[K, V] {
	ks := slices.SortedFunc(maps.Keys(m), compare)
	return func(yield func(K, V) bool) {
		for _, k := range ks {
			v := m[k]
			if !yield(k, v) {
				return
			}
		}
	}
}

// Filter only yields the items where fn returns true.
func Filter[T any](seq iter.Seq[T], fn func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			if fn(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Uniq omits any subsequent items that are equal.
func Uniq[T comparable](seq iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		var v1 T
		v1Set := false
		for v := range seq {
			if v1Set {
				if v == v1 {
					continue
				}
			}
			if !yield(v) {
				return
			}
			v1 = v
			v1Set = true
		}
	}
}

// UniqFunc is like [Uniq], but uses a supplied equal function.
func UniqFunc[T any](seq iter.Seq[T], equal func(T, T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		var v1 T
		v1Set := false
		for v := range seq {
			if v1Set {
				if equal(v, v1) {
					continue
				}
			}
			if !yield(v) {
				return
			}
			v1 = v
			v1Set = true
		}
	}
}

// Join is like [strings.Join], but accepts a seq.
// For example, seq can be the result of a call to [Map].
// Note that this has to do more memory allocations than
// [strings.Join], so it shouldn't be used if you already
// have a slice of strings.
func Join(seq iter.Seq[string], delim string) string {
	var b strings.Builder
	i := 0
	for s := range seq {
		if i != 0 {
			b.WriteString(delim)
		}
		b.WriteString(s)
		i++
	}
	return b.String()
}

// Markovian returns an infinite sequence that is a first-order
// recurrence relation (a sequence where each element only
// depends on the previous element).
// Use e.g. [Take] to limit the number
// of elements.
func Markovian[T any](first T, next func(T) T) iter.Seq[T] {
	return func(yield func(T) bool) {
		a := first
		for {
			if !yield(a) {
				return
			}
			a = next(a)
		}
	}
}

// Take returns the sequence up to at most the
// first n elements.
func Take[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		i := 0
		for v := range seq {
			if i >= n || !yield(v) {
				return
			}
			i++
		}
	}
}

// Wrapper for [fmt.Sprint] that takes one typed parameter.
// Useful for e.g. [Map].
func Sprint[T any](v T) string {
	return fmt.Sprint(v)
}

// All returns true if and only if each element in seq
// is true.
func All(seq iter.Seq[bool]) bool {
	for b := range seq {
		if !b {
			return false
		}
	}
	return true
}

// First keeps only the first item in the [iter.Seq2]
// for each element.
func First[A, B any](seq iter.Seq2[A, B]) iter.Seq[A] {
	return func(yield func(A) bool) {
		for a, _ := range seq {
			if !yield(a) {
				return
			}
		}
	}
}

// First keeps only the second item in the [iter.Seq2]
// for each element.
func Second[A, B any](seq iter.Seq2[A, B]) iter.Seq[B] {
	return func(yield func(B) bool) {
		for _, b := range seq {
			if !yield(b) {
				return
			}
		}
	}
}

// WithFirst sets the first item of each element in the
// resulting [iter.Seq2] to fn applied to the second item.
func WithFirst[A, B any](seq iter.Seq[B], fn func(B) A) iter.Seq2[A, B] {
	return func(yield func(A, B) bool) {
		for b := range seq {
			if !yield(fn(b), b) {
				return
			}
		}
	}
}

// WithSecond sets the second item of each element in the
// resulting [iter.Seq2] to fn applied to the first item.
func WithSecond[A, B any](seq iter.Seq[A], fn func(A) B) iter.Seq2[A, B] {
	return func(yield func(A, B) bool) {
		for a := range seq {
			if !yield(a, fn(a)) {
				return
			}
		}
	}
}

func Repeat[T any](x T, n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		for range n {
			if !yield(x) {
				return
			}
		}
	}
}

// Merge returns an [iter.Seq2] with the elements from
// seqA and seqB. The returned sequence ends when
// any of seqA or seqB runs out of elements
func Merge[A, B any](seqA iter.Seq[A], seqB iter.Seq[B]) iter.Seq2[A, B] {
	return func(yield func(A, B) bool) {
		nextB, stopB := iter.Pull(seqB)
		defer stopB()
		for a := range seqA {
			b, bOk := nextB()
			if !bOk {
				return
			}
			if !yield(a, b) {
				return
			}
		}
	}
}
