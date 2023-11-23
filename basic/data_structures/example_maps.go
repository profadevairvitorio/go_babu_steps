// Maps
// Maps são coleções não ordenadas de pares chave-valor.

package main

import "fmt"

func main() {
	// Declaração e inicialização de um map
	pessoa := map[string]string{
		"nome":  "João",
		"idade": "25",
	}

	// Adicionar um novo par chave-valor
	pessoa["cidade"] = "São Paulo"

	// Acessar um valor usando a chave
	fmt.Println(pessoa["nome"]) // Saída: João

	// Deletar um par chave-valor
	delete(pessoa, "idade")
	fmt.Println(pessoa) // Saída: map[nome:João cidade:São Paulo]
}
