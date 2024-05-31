package main

import (
	"fmt"
	"strconv"
)

func greetings(name string) string {
	return fmt.Sprintf("hello, %s", name)
}

func sum(a, b int) string {
	return fmt.Sprintf("a soma de %d mais %d resulta em %d", a, b, a+b)
}

func convertAndSum(a, b string) (sum string, err error) {
	aConverted, errAConverter := strconv.Atoi(a)

	// tratamento de erro em golang, quando um erro é instanciado é necessário tratar dessa forma
	// validar se ela é diferente de "nil", "nil" corresponde a null no javascript
	if errAConverter != nil {
		return "", fmt.Errorf("erro na conversão")
	}

	// tratamento de erro
	bConverted, errBConverter := strconv.Atoi(b)

	if errBConverter != nil {
		return "", fmt.Errorf("erro na conversão")
	}

	return fmt.Sprintf("Resultado da soma pós conversão: %d", aConverted+bConverted), nil
}

func main() {
	fmt.Println(greetings("Diego"))

	a := 10
	b := 20
	fmt.Println(sum(a, b))

	resultA, _ := convertAndSum("10", "10")
	fmt.Println(resultA)

	_, err := convertAndSum("ç", "10")
	if err != nil {
		fmt.Println(err)
	}
}
