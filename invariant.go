package invar

import (
	"errors"
	"fmt"
)

var (
	ErrInvariantNotHolding = errors.New("invariant does not hold up")
)

type Invariant[T any] struct {
	Name  string
	Check func(T) error
}

func checkValInvariants[T any](val T, invariants []Invariant[T]) error {
	for i, inv := range invariants {
		err := inv.Check(val)
		if err == nil { // all fine
			continue
		}

		invName := inv.Name
		if invName == "" {
			invName = fmt.Sprintf("#%d", i)
		}

		return getInvariantError(invName, err)
	}

	return nil
}

func getInvariantError(invariantName string, err error) error {
	return fmt.Errorf("%w: invariant '%s' fails with error: %w", ErrInvariantNotHolding, invariantName, err)
}
