package tests

import (
	"reflect"
	"testing"
)

type T struct {
	*testing.T
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
	if !reflect.DeepEqual(expect, got) {
		t.Fatalf("Equal(): expected %v, but got %v", expect, got)
	}
}
