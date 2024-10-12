package invar

import (
	"errors"
	"fmt"
)

var (
	ErrInvariantNotHolding = errors.New("invariant does not hold up")
)

type InvariantsHolder[T any] interface {
	get() (T, error)
}

type invariantsHolder[T any] struct {
	internal   T
	invariants []Invariant[T]
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

	return &invariantsHolder[T]{
		internal:   val,
		invariants: invariants,
	}
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

func (ih *invariantsHolder[T]) get() (T, error) {
	err := checkValInvariants(ih.internal, ih.invariants)
	if err != nil {
		var zero T
		return zero, err
	}

	return ih.internal, nil
}

func Unwrap[T any](holder InvariantsHolder[T]) T {
	val, err := holder.get()
	if err != nil {
		panic(err)
	}

	return val
}

func TryUnwrap[T any](holder InvariantsHolder[T]) (T, error) {
	val, err := holder.get()
	if err != nil {
		var zero T
		return zero, fmt.Errorf("error unwrapping invariant holder: %w", err)
	}

	return val, nil
}

func Inited[T any](holder InvariantsHolder[T]) bool {
	return holder != nil
}

func getInvariantError(invariantName string) error {
	return fmt.Errorf("%w ('%s')", ErrInvariantNotHolding, invariantName)
}
