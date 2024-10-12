package person

import (
	"app/pkg"
)

type Person struct {
	Name pkg.NonEmptyString
	Age  pkg.PositiveInt
}

type ValidPerson pkg.Invariant[Person]

func New(p Person) (ValidPerson, error) {
	return pkg.TryNew(p, []pkg.Condition[Person]{
		func(p Person) bool {
			return pkg.Inited(p.Name)
		},
		func(p Person) bool {
			return pkg.Inited(p.Age)
		},
	})
}
