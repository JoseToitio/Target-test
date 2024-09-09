package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "A persistência é o caminho do êxito."
	var choice int

	fmt.Println("Escolha a forma de entrada da string:")
	fmt.Println("1. Entrada do usuário")
	fmt.Println("2. String pré-definida no código")
	fmt.Print("Escolha (1 ou 2): ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Print("Informe uma string: ")
		fmt.Scan(&input)
	} else if choice == 2 {
		fmt.Printf("String pré-definida: %s\n", input)
	} else {
		fmt.Println("Escolha inválida.")
		return
	}

	count := strings.Count(strings.ToLower(input), "a")

	if count > 0 {
		fmt.Printf("A letra 'a' aparece %d vezes na string.\n", count)
	} else {
		fmt.Println("A letra 'a' não aparece na string.")
	}
}
