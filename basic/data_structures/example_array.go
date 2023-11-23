// Arrays
// Um array em Go é uma coleção de elementos do mesmo tipo com um tamanho fixo.

package main

import "fmt"

func main() {
	// Declaração de um array de inteiros com tamanho 5
	var numeros [5]int
	numeros[0] = 1
	numeros[1] = 2
	// ...

	// Inicialização curta
	outroArray := [3]string{"Maçã", "Banana", "Laranja"}

	fmt.Println(numeros)
	fmt.Println(outroArray)
}
