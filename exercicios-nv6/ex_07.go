// - Atribua uma função a uma variável.
// - Chame essa função.
package main

import "fmt"

func main() {
	anony := func() {
		fmt.Println("Funcão_anonima.jpg")
	} // criamos a função e passamos pra uma variavel
	anony() // e chamamos ela depois pela variavel.
}
