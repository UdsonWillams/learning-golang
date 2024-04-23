// - Comece com essa slice:
//     - x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
// - Utilizando slicing e append, crie uma slice y que contenha os valores:
//     - [42, 43, 44, 48, 49, 50, 51]

package main

import "fmt"

// EU FIZ ASSIM, MAS A PROFESSORA FEZ DE UMA FORMA MELHOR.
func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := []int{}

	a := x[:3]
	fmt.Println(a)
	b := x[6:]
	fmt.Println(b)

	y = append(y, a...)
	y = append(y, b...)

	fmt.Println(y)
}

// EXEMPLOS PROFESSORA.
package main

import (
	"fmt"
)

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	
//	x = append(x[:3], x[len(x)-4:]...)
	x = append(x[:3], x[6:]...)	
	
	//[42, 43, 44, 48, 49, 50, 51]
	
	fmt.Println(x)
}
