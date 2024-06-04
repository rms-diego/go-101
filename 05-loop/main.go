package main

import "fmt"

func main() {
	name := []string{"Diego", "Naty", "Tatiana"}

	fmt.Println("FOR CLÁSSICO")
	// for clássico em go
	for i := 0; i < len(name); i++ {
		fmt.Println(name[i])
	}

	fmt.Println("\nFOR RANGE")
	// for usando o range, semelhante ao for of do javascript
	for index, name := range name {
		fmt.Println(index, name)
	}

	fmt.Println("\nWHILE LOOP")
	// while em golang
	index := 0
	for index < len(name) {
		fmt.Println(name[index])
		index++
	}
}
