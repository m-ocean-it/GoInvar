package invar

import (
	"errors"
	"fmt"
)

type InvariantsHolder[T any] interface {
	get() T
}

type invariantsHolder[T any] struct {
	internal T
}

type Invariant[T any] struct {
	Name  string
	Check func(T) bool
}

func New[T any](val T, invariants []Invariant[T]) *invariantsHolder[T] {
	_ = checkValInvariants(val, invariants, true)

	return &invariantsHolder[T]{internal: val}
}

func TryNew[T any](val T, invariants []Invariant[T]) (*invariantsHolder[T], error) {
	err := checkValInvariants(val, invariants, false)
	if err != nil {
		return nil, fmt.Errorf("could not construct invariants holder: %w", err)
	}

	return &invariantsHolder[T]{internal: val}, nil
}

func checkValInvariants[T any](val T, invariants []Invariant[T], panicOnErr bool) error {
	for i, inv := range invariants {
		if inv.Check(val) {
			continue
		}

		invName := inv.Name
		if invName == "" {
			invName = fmt.Sprintf("#%d", i)
		}

		err := fmt.Errorf("invariant '%s' does not hold up", invName)

		if panicOnErr {
			panic(err)
		}

		return err
	}

	return nil
}

func (holder *invariantsHolder[T]) get() T {
	return holder.internal
}

func Unwrap[T any](holder InvariantsHolder[T]) T {
	return holder.get()
}

func TryUnwrap[T any](holder InvariantsHolder[T]) (T, error) {
	if holder == nil {
		var zero T
		return zero, errors.New("invariant was not initialized")
	}

	return holder.get(), nil
}

func Inited[T any](holder InvariantsHolder[T]) bool {
	return holder != nil
}
