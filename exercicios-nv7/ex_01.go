// - Crie um valor e atribua-o a uma variável.
// - Demonstre o endereço deste valor na memória.
package main

import "fmt"

func main() {

	x := 15
	y := "string"

	fmt.Println("ENDEREÇOS NA MEMORIA: ")
	fmt.Println(&x)
	fmt.Println(&y)
	fmt.Println("VALORES CONTIDOS NESSE ENDEREÇO: ")
	a := *&x
	b := *&y
	fmt.Println(a)
	fmt.Println(b)

}
