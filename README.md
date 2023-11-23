## Conceitos Básicos:
### Packages:

O código Go é organizado em pacotes. O pacote principal é main.

### Imports:

O comando import é usado para incluir pacotes externos no seu código.
Funções:

O ponto de entrada do programa é a função main.

### Variáveis:

Declare variáveis usando var ou := para inferência de tipo.

```go
var idade int
nome := "Gopher"
```

### Estrtuturas de controle

Use if, else, for e switch para controle de fluxo.
```go
package main

import "fmt"

func main() {
    idade := 18

    if idade >= 18 {
        fmt.Println("Você é maior de idade.")
    } else {
        fmt.Println("Você é menor de idade.")
    }
}
```

```go
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
```

### Estruturas de repetição
Em Go, a principal estrutura de repetição é o loop for.

```go
// loop simples
package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
}
```

```go
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
```

### Estruturas de Dados:

Arrays, slices, maps e structs são estruturas de dados em Go.

```go
// Arrays
// Um array em Go é uma coleção de elementos do mesmo tipo com um tamanho fixo.

ackage main

import "fmt"

func main() {
// Declaração de um array de inteiros com tamanho 5
var numeros [5]int
numeros[0] = 1
numeros[1] = 2
// ...

// Inicialização curta
outroArray := [3]string{"Maçã", "Banana", "Laranja"}

fmt.Println(numeros)
fmt.Println(outroArray)
}
```

```go
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
```

```go
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

```

```go
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
	fmt.Println(pessoa.Nome) // Saída: Maria
	fmt.Println(pessoa.Idade) // Saída: 30
}
```

### Funções
Em Go, as funções são blocos de código reutilizáveis que realizam uma tarefa específica. 

```go
// função simples
package main

import "fmt"

// Definição de uma função que imprime "Hello, Go!"
func saudacao() {
    fmt.Println("Hello, Go!")
}

func main() {
    // Chamada da função
    saudacao()
}
```
