// // Põe na tela: todos os números de 1 a 10000.

package main

import "fmt"

func main() {
	fmt.Printf("1° formato de for\n")
	for x := 1; x <= 10000; x++ {
		fmt.Printf("%v, ", x)
	}
	fmt.Printf("\n2° formato de for\n")
	x := 1
	for {
		fmt.Printf("%v, ", x)
		x++
		if x > 10000 {
			break
		}
	}
}
