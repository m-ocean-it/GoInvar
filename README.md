# GoInvar: encoding invariants with types

Example:

```go
// positive_int/positive_int.go

package positive_int

import "github.com/m-ocean-it/GoInvar"

type PositiveInt goinvar.Invariant[int]

func New(n int) PositiveInt {
	return goinvar.New(n, []goinvar.Condition[int]{
		func(x int) bool {
			return x > 0
		},
	})
}

func TryNew(n int) (PositiveInt, error) {
	return goinvar.TryNew(n, []goinvar.Condition[int]{
		func(x int) bool {
			return x > 0
		},
	})
}


// main.go

package main

import (
    "github.com/m-ocean-it/GoInvar"

    "app/positive_int"
)

positiveInt := positive_int.New(2)
fmt.Println(goinvar.Unwrap(positiveInt))
```