// - Exercício: tente acessar todos os itens de uma slice sem utilizar range.

package main

import "fmt"

func main() {
	sabores := []string{
		"peperoni", "mozzarela", "abacaxi", "3 queijos",
	}
	len_sabores := len(sabores)
	for i := 0; i < len_sabores; i++ {
		fmt.Println(sabores[i])
	}
}
