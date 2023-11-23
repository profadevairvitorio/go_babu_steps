package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// Criando um leitor de uma string
	reader := strings.NewReader("teste, aqui!")

	// Buffer para armazenar os dados lidos
	buffer := make([]byte, 10)

	// Lendo dados do leitor
	n, err := reader.Read(buffer)
	if err != nil && err != io.EOF {
		fmt.Println("Erro:", err)
		return
	}

	// Imprimindo os dados lidos
	fmt.Printf("Bytes lidos: %d\n", n)
	fmt.Printf("Dados lidos: %s\n", buffer[:n])
}
