package individualPerson

import (
	"fmt"

	"github.com/rms-diego/go-101/09-oque-sao-interface/person"
)

type individualPerson struct {
	person.Person // herda os "atributos" definidas na struct "person"
	cpf           string
}

func NewIndividualPerson(name, lastName, cpf string, age int8) *individualPerson {
	return &individualPerson{
		Person: person.Person{Name: name, LastName: lastName, Age: age},
		cpf:    cpf,
	}
}

func (ip *individualPerson) ShowTaxId() string {
	return fmt.Sprintf("My cpf is: %v", ip.cpf)
}
