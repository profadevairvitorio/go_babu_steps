package main

import (
	"fmt"
	"os"
)

func main() {
	// Abrindo um arquivo para escrita (criando se não existir)
	arquivo, err := os.Create("files/arquivo_os.txt")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}
	defer arquivo.Close()

	// Escrevendo dados no arquivo
	dados := []byte("Conteúdo a ser escrito usando os.")
	_, err = arquivo.Write(dados)
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}

	fmt.Println("Arquivo criado e escrito com sucesso usando os.")
}
