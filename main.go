package main

import (
	"app/person"
	"fmt"
)

func main() {
	p, err := person.New("Simon", 29)
	if err != nil {
		panic(err)
	}

	fmt.Printf("name is %s\n", p.GetName())
	fmt.Printf("age is %d\n", p.GetAge())
}
