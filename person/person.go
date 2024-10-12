package person

import (
	"app/pkg"
	"errors"
)

type person struct {
	name pkg.NonEmptyString
	age  pkg.PositiveInt
}

type ValidPerson interface {
	GetName() string
	GetAge() int

	self() *person
}

func New(name string, age int) (ValidPerson, error) {
	nonEmptyName, err := pkg.TryNewNonEmptyString(name)
	if err != nil {
		return nil, errors.New("")
	}
	positiveAge, err := pkg.TryNewPositiveInt(age)
	if err != nil {
		return nil, errors.New("")
	}

	return &person{
		name: nonEmptyName,
		age:  positiveAge,
	}, nil
}

func (p *person) GetName() string {
	return pkg.Unwrap(p.name)
}

func (p *person) GetAge() int {
	return pkg.Unwrap(p.age)
}

func (p *person) self() *person {
	return p
}
