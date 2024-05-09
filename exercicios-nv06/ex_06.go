// - Crie e utilize uma função anônima.

package main

import "fmt"

func main() {

	x := 10
	func(x int) {
		fmt.Println(x, "vezes 873648 é:")
		fmt.Println(x * 873648)
	}(x) // Criamos a função e chamamos ela na mesma hr

	anony := func() {
		fmt.Println("Funcão_anonima.jpg")
	} // criamos a função e passamos pra uma variavel
	anony() // e chamamos ela depois pela variavel.
	anony() // EM outro momento.
}
