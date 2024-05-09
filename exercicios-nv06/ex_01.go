// - Exercício:
//     - Crie uma função que retorne um int
//     - Crie outra função que retorne um int e uma string
//     - Chame as duas funções
//     - Demonstre seus resultados

package main

import "fmt"

func main() {

	a_int := ReturnInt()
	b_int, c_str := ReturnIntString()
	fmt.Println(a_int)
	fmt.Println(b_int, c_str)
}

func ReturnInt() int {
	fmt.Println("retorna um int")
	return 1
}

func ReturnIntString() (int, string) {
	fmt.Println("retorna um int e uma string")
	return 1, "string"
}
