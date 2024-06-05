package businessperson

import (
	"fmt"

	"github.com/rms-diego/go-101/09-oque-sao-interface/person"
)

type businessperson struct {
	person.Person // herda os "atributos" definidas na struct "person"
	corporateName string
	cnpj          string
}

func NewBusinessperson(name, lastName, corporateName, cnpj string, age int8) *businessperson {
	return &businessperson{
		Person:        person.Person{Name: name, LastName: lastName, Age: age},
		corporateName: corporateName,
		cnpj:          cnpj,
	}
}

func (bp *businessperson) ShowTaxId() string {
	return fmt.Sprintf("My cnpj is: %v\nMy corporate name is: %v", bp.cnpj, bp.corporateName)
}
