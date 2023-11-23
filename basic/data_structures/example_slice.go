// Slices
// Slices são estruturas de dados dinâmicas que representam uma parte de um array.

package main

import "fmt"

func main() {
	// Declaração de uma slice
	frutas := []string{"Maçã", "Banana", "Laranja"}

	// Adicionar um elemento à slice
	frutas = append(frutas, "Morango")

	// Acessar elementos da slice
	fmt.Println(frutas[1]) // Saída: Banana

	// Slicing para obter uma parte da slice
	parte := frutas[1:3]
	fmt.Println(parte) // Saída: [Banana Laranja]
}
