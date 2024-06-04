package main

import "fmt"

type person struct {
	name     string
	lastName string
	age      int8
}

type businessPerson struct {
	person        // herda os "atributos" definidas na struct "person"
	corporateName string
	cnpj          string
}

type individualPerson struct {
	person // herda os "atributos" definidas na struct "person"
	cpf    string
}

// Definindo função greetings como um "método" da struct "person"
func (p person) greetings() string {
	return fmt.Sprintf("My name is %v %v and my age is %v years old", p.name, p.lastName, p.age)
}

func main() {
	individualPerson := individualPerson{
		// ao usar a herança é necessário passar uma pessoa como referencia
		person: person{name: "Diego", lastName: "Ramos", age: 23},
		cpf:    "000.000.000-00",
	}

	businessperson := businessPerson{
		// ao usar a herança é necessário passar uma pessoa como referencia
		person:        person{name: "Tatiana", lastName: "Ramos", age: 48},
		cnpj:          "00.000.000/0001-00",
		corporateName: "Tatiana Ramos Inc",
	}

	fmt.Println(individualPerson.greetings())
	fmt.Println(businessperson.greetings())
}
