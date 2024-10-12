package invar

import (
	"errors"
	"fmt"
)

var ErrCannotUnwrapNilVariantsHolder = errors.New("cannot unwrap a nil invariants holder")

func Unwrap[T any](holder InvariantsHolder[T]) T {
	if holder == nil {
		panic(ErrCannotUnwrapNilVariantsHolder)
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
		return zero, ErrCannotUnwrapNilVariantsHolder
	}

	val, err := holder.get()
	if err != nil {
		var zero T
		return zero, fmt.Errorf("error unwrapping invariant holder: %w", err)
	}

	return val, nil
}
