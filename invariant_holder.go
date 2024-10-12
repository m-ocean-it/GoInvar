package invar

type InvariantsHolder[T any] interface {
	get() (T, error)
	check() error
}

type invariantsHolder[T any] struct {
	internal   T
	invariants []Invariant[T]
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

func (ih *invariantsHolder[T]) get() (T, error) {
	err := ih.check()
	if err != nil {
		var zero T
		return zero, err
	}

	return ih.internal, nil
}

func (ih *invariantsHolder[T]) check() error {
	err := checkValInvariants(ih.internal, ih.invariants)
	if err != nil {
		return err
	}

	return nil
}
