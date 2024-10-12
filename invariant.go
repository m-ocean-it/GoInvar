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
	Check func(T) bool
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

func getInvariantError(invariantName string) error {
	return fmt.Errorf("%w ('%s')", ErrInvariantNotHolding, invariantName)
}
