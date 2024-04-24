package main

import "fmt"

func main() {
	t := somenteImpares(soma, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}...)
	fmt.Println(t)
}

func soma(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}

func somenteImpares(f func(x ...int) int, y ...int) int {
	var slice []int
	for _, v := range y {
		if v%2 != 0 {
			slice = append(slice, v)
		}
	}
	total := f(slice...) // função de callback
	return total
}
