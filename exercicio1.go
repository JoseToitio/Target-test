package main

import "fmt"

func fibonacci(limit int) []int {
	fibSeq := []int{0, 1}
	for {
		next := fibSeq[len(fibSeq)-1] + fibSeq[len(fibSeq)-2]
		if next > limit {
			break
		}
		fibSeq = append(fibSeq, next)
	}
	return fibSeq
}

func isFibonacci(num int, fibSeq []int) bool {
	for _, v := range fibSeq {
		if v == num {
			return true
		}
	}
	return false
}

func main() {
	var num int
	var choice int

	fmt.Println("Escolha a forma de entrada do número:")
	fmt.Println("1. Entrada do usuário")
	fmt.Println("2. Número pré-definido no código")
	fmt.Print("Escolha (1 ou 2): ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Print("Informe um número: ")
		fmt.Scan(&num)
	} else if choice == 2 {
		num = 21
		fmt.Printf("Número pré-definido: %d\n", num)
	} else {
		fmt.Println("Escolha inválida.")
		return
	}

	fibSeq := fibonacci(num)
	if isFibonacci(num, fibSeq) {
		fmt.Printf("O número %d pertence à sequência de Fibonacci.\n", num)
	} else {
		fmt.Printf("O número %d não pertence à sequência de Fibonacci.\n", num)
	}
}
