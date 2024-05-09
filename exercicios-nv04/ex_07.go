// - Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
//   - "Nome", "Sobrenome", "Hobby favorito"
//
// - Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
package main

import "fmt"

func main() {
	slice_mult := [][]string{
		[]string{"udson", "willams", "assistir"},
		[]string{"elida", "rego", "desenhar"},
		[]string{"alonço", "fernandes", "passear"},
	}

	for i, v := range slice_mult {
		for values := range v {
			fmt.Printf("Index: %v e Valores: %v\n", i, values)
		}
	}

}
