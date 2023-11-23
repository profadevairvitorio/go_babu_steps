package main

import (
	"fmt"
	"os"
)

func main() {
	// Criando um escritor para a saída padrão (os.Stdout)
	writer := os.Stdout

	// Escrevendo dados no escritor
	dados := []byte("Hello, Gooooo!")
	n, _ := writer.Write(dados)

	// Imprimindo a quantidade de bytes escritos
	fmt.Printf("Bytes escritos: %d\n", n)
}
