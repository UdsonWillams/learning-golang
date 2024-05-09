// - Crie e use um struct anônimo.
// - Desafio: dentro do struct tenha um
// valor de tipo map e outro do tipo slice.

package main

import "fmt"

func main() {

	anonimo := struct {
		valor_1 map[string]int
		valor_2 []int
	}{
		valor_1: map[string]int{"1": 1, "2": 2},
		valor_2: []int{1, 2, 3, 4, 5},
	}

	fmt.Println(anonimo.valor_1)
	fmt.Println(anonimo.valor_2)
}
