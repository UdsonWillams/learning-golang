// - Põe na tela: O unicode code point de todas as letras maiúsculas do alfabeto, três vezes cada.
// - Por exemplo:
//     65
//         U+0041 'A'
//         U+0041 'A'
//         U+0041 'A'
//     66
//         U+0042 'B'
//         U+0042 'B'
//         U+0042 'B'
//     ...e por aí vai.

package main

import "fmt"

func main() {
	initial_value := 65
	final_value := 90
	fmt.Printf("1° formato de for\n")
	for first_iterator := initial_value; first_iterator <= final_value; first_iterator++ {
		fmt.Printf("\b%v\n", first_iterator)
		for i := 1; i <= 3; i++ {
			fmt.Printf("%#U\n", first_iterator)
		}
	}
	second_iterator := initial_value
	fmt.Printf("\n2° formato de for\n")
	for {
		fmt.Printf("\b%v\n", second_iterator)
		for i := 1; i <= 3; i++ {
			fmt.Printf("%#U\n", second_iterator)
		}
		second_iterator++
		if second_iterator > final_value {
			break
		}
	}
}
