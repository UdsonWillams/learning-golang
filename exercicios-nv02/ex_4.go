// - Crie um programa que:
//   - Atribua um valor int a uma variável
//   - Demonstre este valor em decimal,
//    binário e hexadecimal
//   - Desloque os bits dessa variável 1 para a esquerda,
//   e atribua o resultado a outra variável
//   - Demonstre esta outra variável em decimal,
//   binário e hexadecimal

package main

import "fmt"

func main() {
	var num int = 200
	fmt.Printf("Valor Inteiro: %v\n", num)
	fmt.Printf("Valor Binario: %b\n", num)
	fmt.Printf("Valor Hexadecimal: %#x\n", num)
	var num_deslocado = num << 1
	fmt.Printf("Valor Inteiro: %v\n", num_deslocado)
	fmt.Printf("Valor Binario: %b\n", num_deslocado)
	fmt.Printf("Valor Hexadecimal: %#x\n", num_deslocado)
}