// - Crie uma função que receba um parâmetro variádico do tipo
// int e retorne a soma de todos os ints recebidos;
// - Passe um valor do tipo slice of int como argumento para a função;
// - Crie outra função, esta deve receber um valor do tipo slice of
// int e retornar a soma de todos os elementos da slice;
// - Passe um valor do tipo slice of int como argumento para a função.

package main

import "fmt"

func main() {
	value_1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	total_1 := SomaInts(value_1...)
	fmt.Println(total_1)

	value_2 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	total_2 := SomaSlicesInts(value_2)
	fmt.Println(total_2)

}

func SomaInts(value ...int) int {
	total := 0
	for _, v := range value {
		total += v
	}
	return total
}

func SomaSlicesInts(value []int) int {
	total := SomaInts(value...)
	return total
}
