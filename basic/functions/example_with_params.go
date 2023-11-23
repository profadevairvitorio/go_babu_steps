package main

import "fmt"

// Função que aceita parâmetros
func soma(a, b int) int {
	return a + b
}

func main() {
	// Chamada da função com parâmetros
	resultado := soma(3, 4)
	fmt.Println("3 + 4 =", resultado)
}
