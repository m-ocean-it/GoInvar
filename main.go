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

	n, err := positiveInt.Get()
	if err != nil {
		panic(err)
	}

	fmt.Println(n)

	var x pkg.PositiveInt
	fmt.Println(x)

	y, err := x.Get()
	if err != nil {
		panic(err)
	}

	fmt.Println(y)
}
