package pkg

import (
	"errors"
	"fmt"
)

type Invariant[T any] interface {
	get() T
}

type invariant[T any] struct {
	internal T
}

type Condition[T any] func(T) bool

func NewInvariant[T any](val T, conditions []Condition[T]) (*invariant[T], error) {
	for i, cond := range conditions {
		if !cond(val) {
			return nil, fmt.Errorf("condition #%d does not hold up", i)
		}
	}

	return &invariant[T]{internal: val}, nil
}

func (inv *invariant[T]) get() T {
	return inv.internal
}

func Get[T any](inv Invariant[T]) (T, error) {
	if inv == nil {
		var zero T
		return zero, errors.New("invariant was not initialized")
	}

	return inv.get(), nil
}

func Inited[T any](inv Invariant[T]) bool {
	return inv != nil
}
