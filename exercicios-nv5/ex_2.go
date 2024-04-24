// - Utilizando a solução anterior, coloque os valores do
// tipo "pessoa" num map, utilizando os sobrenomes como key.
// - Demonstre os valores do map utilizando range.
// - Os diferentes sabores devem ser demonstrados utilizando
//  outro range, dentro do range anterior.

package main

import "fmt"

type pessoa struct {
	Nome      string
	Sobrenome string
	sabores   []string
}

func main() {

	pessoa1 := pessoa{
		Nome:      "Udson",
		Sobrenome: "Willams",
		sabores:   []string{"chocolá", "morango"},
	}

	pessoa2 := pessoa{"Elida", "Maria", []string{"balnilha", "morango"}}

	fmt.Print(pessoa1, pessoa2)

	map_pessoa := map[string]pessoa{
		pessoa1.Sobrenome: pessoa1,
		pessoa2.Sobrenome: pessoa2,
	}

	for k, v := range map_pessoa {
		fmt.Printf("\nchave: %v | com pessoa: %v", k, v)
		for _, value := range v.sabores {
			fmt.Printf("\nsabores preferidos: %v", value)
		}
	}
	fmt.Println()
}
