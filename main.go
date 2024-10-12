package main

import (
	"app/person"
	"app/pkg"
	"fmt"
)

func main() {
	p, err := person.New(person.Person{
		Name: pkg.NewNonEmptyString("Oleg"),
		Age:  pkg.NewPositiveInt(36),
	})
	if err != nil {
		panic(err)
	}

	originalPerson := pkg.Unwrap(p)

	fmt.Println(pkg.Unwrap(originalPerson.Name))
}
