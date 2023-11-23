package main

import "fmt"

// Função que retorna múltiplos valores
func divideEresto(dividendo, divisor int) (int, int) {
	quociente := dividendo / divisor
	resto := dividendo % divisor
	return quociente, resto
}

func main() {
	// Chamada da função com múltiplos retornos
	q, r := divideEresto(10, 3)
	fmt.Printf("Quociente: %d, Resto: %d\n", q, r)
}
