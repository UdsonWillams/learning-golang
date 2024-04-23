// - Usando uma literal composta:
//      - Crie um array que suporte 5 valores to tipo int
//      - Atribua valores aos seus índices
// - Utilize range e demonstre os valores do array.
// - Utilizando format printing, demonstre o tipo do array.
package main

import "fmt"

var one_array = [5]int{
	1, 2, 3, 4, 5,
}

func main() {

	for i, v := range one_array {
		fmt.Printf("No index: %v o valor: %v\n", i, v)
	}
	fmt.Printf("tipo do array: %T", one_array)
}
