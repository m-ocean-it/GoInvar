package invar

import (
	"errors"
	"fmt"
)

var ErrNilInvariantHolder = errors.New("invariant holder is nil")

func Unwrap[T any](holder InvariantsHolder[T]) T {
	if holder == nil {
		panic(ErrNilInvariantHolder)
	}

	val, err := holder.get()
	if err != nil {
		panic(err)
	}

	return val
}

func TryUnwrap[T any](holder InvariantsHolder[T]) (T, error) {
	if holder == nil {
		var zero T
		return zero, ErrNilInvariantHolder
	}

	val, err := holder.get()
	if err != nil {
		var zero T
		return zero, fmt.Errorf("error unwrapping invariant holder: %w", err)
	}

	return val, nil
}

func Check[T any](holder InvariantsHolder[T]) error {
	if holder == nil {
		return ErrNilInvariantHolder
	}

	return holder.check()
}

func NamedCheck[T any](name string, holder InvariantsHolder[T]) error {
	if err := Check(holder); err != nil {
		return fmt.Errorf("check '%s' failed: %w", name, err)
	}

	return nil
}
