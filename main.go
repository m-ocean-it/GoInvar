package main

import (
	"app/pkg"
	"fmt"
)

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

	nes, err := pkg.NewNonEmptyString("hello")
	if err != nil {
		panic(err)
	}

	s, err := pkg.Get(nes)
	if err != nil {
		panic(err)
	}

	fmt.Println(s)
}
