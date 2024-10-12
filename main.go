package main

import (
	"fmt"
	"github.com/m-ocean-it/GoInvar/person"
	"github.com/m-ocean-it/GoInvar/pkg"
)

func main() {
	wrappedPerson, err := person.New("Simon", 29)
	if err != nil {
		panic(err)
	}

	validPerson := pkg.Unwrap(wrappedPerson)

	validName := pkg.Unwrap(validPerson.Name)
	validAge := pkg.Unwrap(validPerson.Age)

	fmt.Printf("name is %s\n", validName)
	fmt.Printf("age is %d\n", validAge)
}
