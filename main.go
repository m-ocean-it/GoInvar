package main

import (
	"app/pkg"
	"fmt"
)

type data struct {
	n pkg.PositiveInt
}

func main() {
	positiveInt, err := pkg.NewPositiveInt(2)
	if err != nil {
		panic(err)
	}

	n, err := pkg.Get(positiveInt)
	if err != nil {
		panic(err)
	}

	fmt.Println(n)

	var x pkg.PositiveInt
	x, err = pkg.NewPositiveInt(3)
	if err != nil {
		panic(err)
	}

	y, err := pkg.Get(x)
	if err != nil {
		panic(err)
	}

	fmt.Println(y)
}
