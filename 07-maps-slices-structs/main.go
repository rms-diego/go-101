package main

import "fmt"

func main() {
	// slices
	names := []string{"Diego", "Tatiana", "Nathália"}

	fmt.Println(names)

	names = append(names, "Lucas")
	fmt.Println(names)

	names = append(names, "Aoba")
	fmt.Println(names)

	// maps
	ages := make(map[string]uint8)

	ages["Diego"] = 23
	ages["Nathália"] = 21
	ages["Tatiana"] = 48

	fmt.Println(ages)

	// struct
	type person struct {
		name     string
		lastName string
		age      uint8
		status   bool
	}

	firstPerson := person{
		name:     "Diego",
		lastName: "Ramos",
		age:      23,
		status:   true,
	}

	fmt.Println(firstPerson)
	fmt.Println(firstPerson.name + " " + firstPerson.lastName)
}
