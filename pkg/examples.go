package pkg

type NegativeInt Invariant[int]

func NewNegativeInt(n int) (NegativeInt, error) {
	return NewInvariant(n, []Condition[int]{
		func(x int) bool {
			return x < 0
		},
	})
}

type PositiveInt Invariant[int]

func NewPositiveInt(n int) (PositiveInt, error) {
	return NewInvariant(n, []Condition[int]{
		func(x int) bool {
			return x >= 0
		},
	})
}

type NonEmptyString Invariant[string]

func NewNonEmptyString(s string) (NonEmptyString, error) {
	return NewInvariant(s, []Condition[string]{
		func(s string) bool {
			return len(s) > 0
		},
	})
}
