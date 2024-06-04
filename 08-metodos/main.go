package main

import "fmt"

type Person struct {
	name       string
	lastName   string
	age        int8
	isMarriage bool
}

// Essa função seria o mais proximo ao um "constructor" em go
func newPerson(name, lastName string, age int8, isMarriage bool) (Person, error) {
	if age < 0 {
		return Person{}, fmt.Errorf("age must be greater than 0")
	}

	return Person{
		name,
		lastName,
		age,
		isMarriage,
	}, nil
}

// O "()" antes do nome da função é uma funcionalidade chamada "Receptor"
// O receptor associa uma função a um struct, é possível receber o struct
// Como se fosse um "this", só que o "this" fica na variável do receptor
// No exemplo abaixo o "this" é a variável "p"
func (p *Person) greetings() string {
	isMarriage := func(isMarriage bool) string {
		if isMarriage {
			return "married"
		}

		return "not married"
	}(p.isMarriage)

	return fmt.Sprintf("Hello my name is %v %v\nMy age is %v years old\nAnd i'm a %v",
		p.name,
		p.lastName,
		p.age,
		isMarriage,
	)
}

func main() {

	person1, err := newPerson("Diego", "Ramos", 23, true)

	if err != nil {
		panic(err)
	}

	person2, err := newPerson("Marcio", "Apelão", 89, false)

	if err != nil {
		panic(err)
	}

	fmt.Println(person1.greetings())
	fmt.Println("\n" + person2.greetings())
}
