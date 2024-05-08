// Faça esse código funcionar: https://play.golang.org/p/oB-p3KMiH6
package main

import (
	"fmt"
)

// Deixando o canal bidimensional
func main() {
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}

// Também é possivel excluir o print(<-cs)
// Esse exemplo dá um canal com apenas uma dimensão
// Que recebe dados, mas não os retorna. 

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	cs := make(chan<- int)

// 	go func() {
// 		cs <- 42
// 	}()
// 	fmt.Println(<-cs)

// 	fmt.Printf("------\n")
// 	fmt.Printf("cs\t%T\n", cs)
// }
