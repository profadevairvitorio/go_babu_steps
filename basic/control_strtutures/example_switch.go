package main

import "fmt"

func main() {
	diaDaSemana := "quarta"

	switch diaDaSemana {
	case "segunda", "terça", "quarta", "quinta", "sexta":
		fmt.Println("Dia útil.")
	case "sábado", "domingo":
		fmt.Println("Fim de semana.")
	default:
		fmt.Println("Dia inválido.")
	}
}
