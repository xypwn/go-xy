package ds

import "iter"

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return Set[T]{}
}

// Has returns whether s contains x.
func (s Set[T]) Has(x T) bool {
	_, ok := s[x]
	return ok
}

// Add inserts all elements x into
// the set.
func (s Set[T]) Add(x ...T) {
	for _, x := range x {
		s[x] = struct{}{}
	}
}

// Sub deletes all elements in s
// also contained in other.
func (s Set[T]) Sub(other Set[T]) {
	for k := range other {
		delete(s, k)
	}
}

// Union inserts all elements contained
// in other.
func (s Set[T]) Union(other Set[T]) {
	for k := range other {
		s.Add(k)
	}
}

// Clone returns a copy of s.
func (s Set[T]) Clone() Set[T] {
	res := make(Set[T], len(s))
	for k := range s {
		res.Add(k)
	}
	return res
}

// Values returns an iterator over all keys
// in the set. The order is undefined.
func (s Set[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for k := range s {
			if !yield(k) {
				return
			}
		}
	}
}
