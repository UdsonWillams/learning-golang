// - Crie uma função que retorna uma função.
// - Atribua a função retornada a uma variável.
// - Chame a função retornada.
package main

import "fmt"

func main() {
	faz_soma := RetornaFuncSoma()

	valor_1 := 10
	valor_2 := 5

	total := faz_soma(valor_1, valor_2)
	fmt.Println(total)

}

func RetornaFuncSoma() func(int, int) int {
	return func(x int, y int) int {
		return x + y
	}
}
