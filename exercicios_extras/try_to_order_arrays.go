"""
Crie um array referente aos anos de nascimento das 5 pessoas mais próximas a você. Em seguida:
- Ordene o array  na ordem crescente e mostre o resultado;
- Ordene o array na ordem decrescente e mostre o resultado;
"""

package main

import (
	"fmt"
	"sort"
)

// Exemplo retornado na net.
func sortSliceAscDesc(slice_to_order []int, order string) []int {
	if order == "desc" {
		sort.Slice(slice_to_order, func(i, j int) bool {
			// retorno dá função anonima
			return slice_to_order[i] > slice_to_order[j]
		})
		return slice_to_order
	}
	sort.Slice(slice_to_order, func(i, j int) bool {
		// retorno dá função anonima
		return slice_to_order[i] < slice_to_order[j]
	})
	return slice_to_order
}

func main() {

	// definindo o slice.
	slice := []int{3, 2, 6, 7, 10}
	//second_slice := []int{}

	slice = sortSliceAscDesc(slice, "asc")
	// mostrando o resultado.
	fmt.Println("-----ascendente----")
	for _, v := range slice {
		fmt.Println(v)
	}

	slice = sortSliceAscDesc(slice, "desc")
	fmt.Println("-----descendente----")
	// mostrando o resultado.
	for _, v := range slice {
		fmt.Println(v)
	}
}

// exemplo que peguei na internet e tentei incremetar.
// https://stackoverflow.com/questions/37695209/golang-sort-slice-ascending-or-descending

// --- formato mais simples possivel ---

// a := []int{5, 3, 4, 7, 8, 9}
// sort.Slice(a, func(i, j int) bool {
//     return a[i] < a[j]
// })
// for _, v := range a {
//     fmt.Println(v)
// }
