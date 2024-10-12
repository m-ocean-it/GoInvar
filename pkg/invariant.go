package pkg

import "errors"

type Invariant[T any] interface {
	Get() (T, error)
	self() *invariant[T]
}

type invariant[T any] struct {
	internal T
}

func NewInvariant[T any](val T, conditions []func(T) bool) (*invariant[T], error) {
	for _, cond := range conditions {
		if !cond(val) {
			return nil, errors.New("wrong val")
		}
	}

	return &invariant[T]{internal: val}, nil
}

func (inv *invariant[T]) Get() (T, error) {
	if inv == nil {
		var zero T
		return zero, errors.New("unintitialized invariant")
	}

	return inv.internal, nil
}

func (inv *invariant[T]) self() *invariant[T] {
	return inv
}
