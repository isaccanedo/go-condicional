package main

// @author: Isac Canedo

import (
	"fmt"
	"strings"
)

func main() {
	names()
}

func names() {
	fmt.Println("Digite um nome:")

	var name string
	fmt.Scanln(&name)
	// Verifica se o nome tem vogal
	for _, v := range strings.ToLower(name) {
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
			fmt.Println("Seu nome contém uma vogal!")
			return
		}
	}
	fmt.Println("Seu nome não contém vogal!")
}
