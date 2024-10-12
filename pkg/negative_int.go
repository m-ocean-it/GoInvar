package pkg

func NewNegativeInt(n int) (Invariant[int], error) {
	return NewInvariant(n, []func(int) bool{
		func(x int) bool {
			return x < 0
		},
	})
}
