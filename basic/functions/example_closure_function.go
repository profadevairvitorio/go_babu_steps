package main

import "fmt"

func main() {
	// Função anônima
	quadrado := func(x int) int {
		return x * x
	}

	// Chamada da função anônima
	resultado := quadrado(5)
	fmt.Println("Quadrado de 5:", resultado)
}
