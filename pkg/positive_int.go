package pkg

import "errors"

type PositiveInt interface {
	Int() (int, error)
	self() *positiveInt
}

func NewPositiveInt(n int) (PositiveInt, error) {
	if n <= 0 {
		return nil, errors.New("non-positive integer")
	}

	return &positiveInt{internal: n}, nil
}

type positiveInt struct {
	internal int
}

func (pi *positiveInt) Int() (int, error) {
	if pi == nil {
		return 0, errors.New("positive int is uninitialized")
	}
	return pi.internal, nil
}

func (pi *positiveInt) self() *positiveInt {
	return pi
}
