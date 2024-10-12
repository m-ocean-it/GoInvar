package goinvar

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

func New[T any](val T, conditions []Condition[T]) *invariant[T] {
	for i, cond := range conditions {
		if !cond(val) {
			panic(fmt.Errorf("condition #%d does not hold up", i))
		}
	}

	return &invariant[T]{internal: val}
}

func TryNew[T any](val T, conditions []Condition[T]) (*invariant[T], error) {
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

func Unwrap[T any](inv Invariant[T]) T {
	return inv.get()
}

func TryUnwrap[T any](inv Invariant[T]) (T, error) {
	if inv == nil {
		var zero T
		return zero, errors.New("invariant was not initialized")
	}

	return inv.get(), nil
}

func Inited[T any](inv Invariant[T]) bool {
	return inv != nil
}
