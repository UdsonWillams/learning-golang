// - Escreva expressões utilizando os operadores,
// e atribua seus valores a variáveis.
// ==
// !=
// <=
// <
// >=
// >
// - Demonstre estes valores.

package main

import "fmt"

const val = iota

func main() {
	eq := 1 == 1      // true
	dif := 1 != 1     // false
	less_eq := 1 <= 1 // true
	less := 1 < 2     // true
	gt_eq := 1 >= 0   // true
	gt := 1 > 0 == true

	fmt.Printf("Valor do 1°: %v\n", eq)
	fmt.Printf("Valor do 2°: %v\n", dif)
	fmt.Printf("Valor do 3°: %v\n", less_eq)
	fmt.Printf("Valor do 4°: %v\n", less)
	fmt.Printf("Valor do 5°: %v\n", gt_eq)
	fmt.Printf("Valor do 6°: %v\n", gt)
}
