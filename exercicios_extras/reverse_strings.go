// Faça uma função/método que receba uma string como parâmetro e que
// retorne uma nova string, onde a sequência dos caracteres foi invertida.
// Dentro da parte principal (main), leia uma string digitada pelo usuário e
// passe para a função/método criada, imprimindo em seguida a string devolvida.

package main

import "fmt"

func invert_str(value string) string {
	runes := []rune(value)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {

	var value string
	fmt.Println("Informe o valor a ser invertido: ")
	fmt.Scanln(&value)

	response := invert_str(value)
	fmt.Printf("Você passou o valor: %v\nValor invertido: %v\n", value, response)
}
