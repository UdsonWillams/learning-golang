// - Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
//     - Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
//     - Do quinto ao último item do slice (incluindo o último item!)
//     - Do segundo ao sétimo item do slice (incluindo o sétimo item!)
//     - Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
//     - Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item

package main

import "fmt"

func main() {

	one_array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	one_array_lenght := len(one_array)
	fmt.Println(one_array[:3])
	fmt.Println(one_array[4:])
	fmt.Println(one_array[1:7])
	fmt.Println(one_array[2 : one_array_lenght-1])
}
