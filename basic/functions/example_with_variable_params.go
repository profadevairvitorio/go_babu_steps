package main

import "fmt"

// Função com parâmetros variádicos
func somaVariadica(numeros ...int) int {
	resultado := 0
	for _, num := range numeros {
		resultado += num
	}
	return resultado
}

func main() {
	// Chamada da função com número variável de argumentos
	total := somaVariadica(1, 2, 3, 4, 5, 6)
	fmt.Println("Soma:", total)
}
