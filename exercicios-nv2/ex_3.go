// - Crie constantes tipadas e não-tipadas.
// - Demonstre seus valores.

package main

import "fmt"

const val_1 string = "olá"
const val_2 = 10

func main() {
	fmt.Printf("Valor do 1°: %T\n", val_1)
	fmt.Printf("Valor do 2°: %T\n", val_2)
}
