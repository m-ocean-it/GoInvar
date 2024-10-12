package person

import (
	"errors"

	"github.com/m-ocean-it/GoInvar/pkg"
)

type Person struct {
	Name pkg.NonEmptyString
	Age  pkg.PositiveInt
}

type ValidPerson pkg.Invariant[Person]

func New(name string, age int) (ValidPerson, error) {
	nonEmptyName, err := pkg.TryNewNonEmptyString(name)
	if err != nil {
		return nil, errors.New("name is invalid")
	}
	positiveAge, err := pkg.TryNewPositiveInt(age)
	if err != nil {
		return nil, errors.New("age is invalid")
	}

	p := Person{Name: nonEmptyName, Age: positiveAge}

	return pkg.TryNew(p, []pkg.Condition[Person]{
		func(p Person) bool { return pkg.Inited(p.Name) },
		func(p Person) bool { return pkg.Inited(p.Age) },
	})
}
