package pkg

type NegativeInt Invariant[int]

func NewNegativeInt(n int) (NegativeInt, error) {
	return TryNew(n, []Condition[int]{
		func(x int) bool {
			return x < 0
		},
	})
}

type PositiveInt Invariant[int]

func NewPositiveInt(n int) PositiveInt {
	return New(n, []Condition[int]{
		func(x int) bool {
			return x >= 0
		},
	})
}

func TryNewPositiveInt(n int) (PositiveInt, error) {
	return TryNew(n, []Condition[int]{
		func(x int) bool {
			return x >= 0
		},
	})
}

type NonEmptyString Invariant[string]

func TryNewNonEmptyString(s string) (NonEmptyString, error) {
	return TryNew(s, []Condition[string]{
		func(s string) bool {
			return len(s) > 0
		},
	})
}

func NewNonEmptyString(s string) NonEmptyString {
	return New(s, []Condition[string]{
		func(s string) bool {
			return len(s) > 0
		},
	})
}
