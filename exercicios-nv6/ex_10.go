// - Demonstre o funcionamento de um closure.
// - Ou seja: crie uma função que retorna outra função,
// onde esta outra função faz uso de uma variável alem
// de seu scope.

package main

import "fmt"

func main() {
	value := Closure()
	fmt.Println("ultimo valor: ", value())
}

func Closure() func() int {
	x := 1
	return func() int {
		for i := 0; i <= 20; i++ {
			if i%2 == 0 {
				fmt.Println("valor par: ", i)
				x = i
			}
		}
		return x
	}
}
