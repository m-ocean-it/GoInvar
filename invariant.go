package invar

import (
	"errors"
	"fmt"
)

var (
	ErrInvariantNotHolding = errors.New("invariant not satisfied")
)

type Invariant[T any] func(T) error

func checkValInvariants[T any](val T, invariants []Invariant[T]) error {
	for _, inv := range invariants {
		err := inv(val)
		if err == nil { // all fine
			continue
		}

		return fmt.Errorf("%w: %w", ErrInvariantNotHolding, err)
	}

	return nil
}
