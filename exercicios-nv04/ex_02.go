// - Usando uma literal composta:
//   - Crie uma slice de tipo int
//   - Atribua 10 valores a ela
//
// - Utilize range para demonstrar todos estes valores.
// - E utilize format printing para demonstrar seu tipo.
package main

import "fmt"

func main() {

	one_slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i, v := range one_slice {
		fmt.Printf("No index: %v o valor: %v\n", i, v)
	}
	fmt.Printf("tipo da variavel: %T", one_slice)
}
