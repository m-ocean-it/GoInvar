package invar

import (
	"errors"
	"fmt"
)

var (
	ErrInvariantNotHolding = errors.New("invariant does not hold up")
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
	err := checkValInvariants(val, invariants)
	if err != nil {
		panic(err)
	}

	return &invariantsHolder[T]{internal: val}
}

func TryNew[T any](val T, invariants []Invariant[T]) (*invariantsHolder[T], error) {
	err := checkValInvariants(val, invariants)
	if err != nil {
		return nil, err
	}

	return &invariantsHolder[T]{internal: val}, nil
}

func checkValInvariants[T any](val T, invariants []Invariant[T]) error {
	for i, inv := range invariants {
		if inv.Check(val) {
			continue
		}

		invName := inv.Name
		if invName == "" {
			invName = fmt.Sprintf("#%d", i)
		}

		return getInvariantError(invName)
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

func getInvariantError(invariantName string) error {
	return fmt.Errorf("%w ('%s')", ErrInvariantNotHolding, invariantName)
}
