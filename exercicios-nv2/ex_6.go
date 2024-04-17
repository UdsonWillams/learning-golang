// - Utilizando iota, crie 4 constantes
// cujos valores sejam os próximos 4 anos.
// - Demonstre estes valores.

package main

import "fmt"

const (
	_ = iota + 2024
	years_25
	years_26
	years_27
	years_28
)

func main() {
	fmt.Println(years_25)
	fmt.Println(years_26)
	fmt.Println(years_27)
	fmt.Println(years_28)
}
