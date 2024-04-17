package main

import "fmt"

func main() {
	num_1 := 200
	num_2 := 2
	num_mult := num_1 * num_2
	num_div := num_1 / num_2

	fmt.Printf("Mult:%v\n", num_mult)
	fmt.Printf("Divisão:%v\n", num_div)

	num_mult = num_1 << 1
	num_div = num_1 >> 1
	fmt.Printf("Mult:%v\n", num_mult)
	fmt.Printf("Divisão:%v", num_div)
}
// Output
// >> Mult:400
// >> Divisão:100
// >> Mult:400
// >> Divisão:100