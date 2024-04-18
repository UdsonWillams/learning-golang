// - Crie um programa que utilize a declaração switch, sem switch statement especificado.
// - Utilizando a solução anterior, adicione as opções else if e else.

package main

import "fmt"

func main() {
	tired := 2
	switch {
	case tired == 0:
		fmt.Println("buxim chei")
	case tired == 1:
		fmt.Println("buxim seco")
	case tired == 2:
		fmt.Println("oppaa")
	}
}
