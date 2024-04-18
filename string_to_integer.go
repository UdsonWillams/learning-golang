// Treinando um pouico de slice
// e fazer uma conversão de int para str.
// o padrão string(int) retorna o valor do int na tabela ascii.
package main

import (
	"fmt"
	"strconv"
)

func main() {

	// ItoA - Integer to ASCII
	var slice_to_append = []string{}
	for x := 1; x <= 100; x++ {
		slice_to_append = append(slice_to_append, strconv.Itoa(x))
	}
	fmt.Print("\n", slice_to_append, "\n")
	for index, value := range slice_to_append {
		fmt.Printf("\nO numero: %v < tem valor ascii: %v <", value, string(index+1))
	}
}
