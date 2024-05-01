package main

import (
	"fmt"
	"sort"
)

func main() {
	n1 := []int{1, 2, 3, 0, 0, 0}
	n2 := []int{2, 5, 6}
	m := 3
	n := 3
	fmt.Println(merge(n1, m, n2, n))
}

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	new_slice := []int{}
	for _, v := range nums1 {
		if v != 0 {
			new_slice = append(new_slice, v)
		}

	}
	for _, v := range nums2 {
		if v != 0 {
			new_slice = append(new_slice, v)
		}

	}
	sort.Slice(new_slice, func(i, j int) bool {
		// retorno dá função anonima
		return new_slice[i] < new_slice[j]
	})
	return new_slice
}

// O melhor retorno pro leetcode é esse, mas não fez sentido pra mim o jeito
// que eles queriam o retorno, sendo diretamente no nums2 e não aceitando
// nem que nums1 receba outros valores.
// func merge(nums1 []int, m int, nums2 []int, n int)  {
//     for n > 0 {
//         nums1[m] = nums2[n-1]
//         m++
//         n--
//     }
//     sort.Ints(nums1)
// }
