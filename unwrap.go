package invar

import (
	"fmt"
)

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
