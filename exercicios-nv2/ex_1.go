// - Escreva um programa que mostre um número em decimal, 
// binário, e hexadecimal.
package main

import "fmt"

var num = 10

func main() {
	fmt.Printf("Valor de %v em:\n", num)
	fmt.Printf(" - decimal: %v\n", num)
	fmt.Printf(" - binario: %b\n", num)
	fmt.Printf(" - hexadecimal: %#x\n", num)
}
