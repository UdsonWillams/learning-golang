// - Crie um tipo "pessoa" com tipo subjacente struct,
// que possa conter os seguintes campos:
//     - Nome
//     - Sobrenome
//     - Sabores favoritos de sorvete
// - Crie dois valores do tipo "pessoa" e demonstre estes
// valores, utilizando range na slice que contem os sabores
//  de sorvete.

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

	for i, v := range pessoa1.sabores {
		fmt.Printf("\nindex: %v e sorvete de tipo: %v", i, v)
	}
}
