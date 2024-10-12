package person

import "app/pkg"

type person struct {
	name pkg.NonEmptyString
	age  pkg.PositiveInt
}

type Person pkg.Invariant[person]

func NewPerson(p person) (Person, error) {
	return pkg.NewInvariant(p, []pkg.Condition[person]{
		func(p person) bool {
			return pkg.Inited(p.name)
		},
		func(p person) bool {
			return pkg.Inited(p.age)
		},
	})
}
