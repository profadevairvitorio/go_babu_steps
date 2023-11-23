package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Abrindo um arquivo para leitura
	arquivoEntrada, err := os.Open("files/entrada.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de entrada:", err)
		return
	}
	defer arquivoEntrada.Close()

	// Criando um arquivo para escrita
	arquivoSaida, err := os.Create("files/saida.txt")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo de saída:", err)
		return
	}
	defer arquivoSaida.Close()

	// Copiando os dados do leitor para o escritor
	quantidadeCopiada, err := io.Copy(arquivoSaida, arquivoEntrada)
	if err != nil {
		fmt.Println("Erro ao copiar os dados:", err)
		return
	}

	fmt.Printf("Quantidade de bytes copiados: %d\n", quantidadeCopiada)
}
