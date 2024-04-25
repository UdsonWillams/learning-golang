// - Callback: passe uma função como argumento a outra função.
package main

import "fmt"

func main() {

	valor_1 := 10
	valor_2 := 5
	SomaPlusSubtracao(Subtracao, valor_1, valor_2)

}

func Subtracao(v1 int, v2 int) int {
	return v1 - v2
}

func SomaPlusSubtracao(f func(int, int) int, v1 int, v2 int) {
	fmt.Println("A SOMA É:", v1+v2)
	fmt.Println("E a subtração é:", f(v1, v2))
}
