// - Desafio surpresa!
// - Format printing:
//     - Decimal       %d
//     - Hexadecimal   %#x
//     - Unicode       %#U
//     - Tab           \t
//     - Linha nova    \n
// - Faça um loop dos números 33 a 122, e utilize format
// printing para demonstrá-los como o seu valor em texto/string.
// relacionado a tabela ascii.

package main

import "fmt"

func main() {
	for x := 32; x <= 122; x++ {
		fmt.Printf("%c - \n", x)
	}
}
