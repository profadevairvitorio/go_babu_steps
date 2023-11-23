// loop infinito
package main

import "fmt"

func main() {
	contador := 0

	for {
		fmt.Println("Este é um loop infinito.")
		contador++

		if contador == 3 {
			break
		}
	}
}
