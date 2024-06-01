package main

import (
	"fmt"
	"os"
)

func main() {
	a, b := 13, 13

	if a > b {
		fmt.Println("A é maior do que B")
	} else if a < b {
		fmt.Println("B é maior do que A")
	} else {
		fmt.Println("Os 2 são iguais")
	}

	switch {
	case a > b:
		fmt.Println("A é maior do que B")

	case a < b:
		fmt.Println("B é maior do que A")

	default:
		fmt.Println("Os 2 são iguais")
	}

	file, err := os.Open("./hello.txt")

	if err != nil {
		panic(err)
	}

	data := make([]byte, 100)

	if _, err := file.Read(data); err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
