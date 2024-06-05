package main

import (
	"fmt"

	"github.com/rms-diego/go-101/09-oque-sao-interface/businessperson"
	"github.com/rms-diego/go-101/09-oque-sao-interface/individualPerson"
)

func main() {
	p := individualPerson.NewIndividualPerson("Diego", "Ramos", "000.000.000-00", 23)
	bp := businessperson.NewBusinessperson("Tatiana", "Ramos", "Tatiana inc", "aoba", 48)

	fmt.Println(p.ShowTaxId())
	fmt.Println(bp.ShowTaxId())
}
