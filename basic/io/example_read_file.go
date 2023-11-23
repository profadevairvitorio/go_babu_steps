package main

import (
	"fmt"
	"os"
)

func main() {
	// Lendo todo o conteúdo de um arquivo usando os.ReadFile
	dados, err := os.ReadFile("files/exemplo.txt")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	// Convertendo os dados lidos para uma string e imprimindo
	fmt.Println("Conteúdo do arquivo:")
	fmt.Println(string(dados))
}
