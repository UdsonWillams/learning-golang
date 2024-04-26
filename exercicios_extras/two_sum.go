// Dado um array de números inteiros e um alvo inteiro,
// retorna índices dos dois números de forma que eles somem ao alvo.

// Você pode assumir que cada entrada teria exatamente uma solução,
// e você não pode usar o mesmo elemento duas vezes.

// Você pode retornar a resposta em qualquer ordem.
package main

import "fmt"

func main() {
	value_1 := []int{2, 7, 11, 15}
	value_2 := []int{3, 2, 4}
	value_3 := []int{3, 3}

	fmt.Println(twoSum(value_1, 9))
	fmt.Println(twoSum(value_2, 6))
	fmt.Println(twoSum(value_3, 6))

}

func twoSum(nums []int, target int) []int {
	step := 1
	for index, number := range nums {
		for idx, nw := range nums[step:] {
			if number+nw == target {
				return []int{index, idx + step}
			}
		}
		step += 1
	}
	return []int{}
}
