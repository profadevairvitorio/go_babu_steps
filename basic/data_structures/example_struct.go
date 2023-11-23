// Structs
// Structs são estruturas de dados compostas que podem conter campos com diferentes tipos.

package main

import "fmt"

// Definição de uma struct
type Pessoa struct {
	Nome  string
	Idade int
}

func main() {
	// Inicialização de uma struct
	pessoa := Pessoa{"Maria", 30}

	// Acessar campos da struct
	fmt.Println(pessoa.Nome)  // Saída: Maria
	fmt.Println(pessoa.Idade) // Saída: 30
}
