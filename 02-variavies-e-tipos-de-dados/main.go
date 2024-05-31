package main

import (
	"fmt"
	"time"
)

// Variáveis a nível de package
var name string

func main() {
	name = "Diego Ramos"
	message := fmt.Sprintf("olá, %s tudo bem?", name) // interpolação de string

	fmt.Println(message)

	// declarando múltiplas variáveis de tipo diferente
	var (
		age        int8      = 10
		birth_date time.Time = time.Now()
	)

	fmt.Println(age)
	fmt.Println(birth_date)

	// declarando múltiplas variáveis do mesmo tipo
	var a, b, c = 1, 2, 3
	fmt.Println(a, b, c)
}
