// - Nível 10?! Êita! Parabéns!
// - Faça esse código funcionar: https://play.golang.org/p/j-EA6003P0
//     - Usando uma função anônima auto-executável
//     - Usando buffer
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() { c <- 42 }()

	fmt.Println(<-c)
}
// Utilizando função anonima.

// Também é possivel usar o buffer na hora de setar o canall
// Ex.:
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	c := make(chan int, 1)

// 	c <- 42

// 	fmt.Println(<-c)
// }
