/*
package calculadora faz calculos matematicos
*/
package calculadora

import "fmt"

// X é o PRIMEIRO valor
var x int

// Y é o SEGUNDO valor
var y int

// Func main Retorna os valores somados, subtraidos, divididos e multiplicados.
func main() {
	x = 10
	y = 25
	fmt.Println(Somar(x, y))
	fmt.Println(Sub(x, y))
	fmt.Println(div(x, y))
	fmt.Println(mult(x, y))
}

// Func somar soma o primeiro e segundo valor
func Somar(x, y int) int {
	return x + y
}

// Func sub subtrai o primeiro e segundo valor
func Sub(x, y int) int {
	return x - y
}

// Func div divide o primeiro e segundo valor
// NÃO VAI APARECER POR ESTAR EM MINUSCULO
func div(x, y int) int {
	return x / y
}

// Func mult multiplica o primeiro e segundo valor
// NÃO VAI APARECER POR ESTAR EM MINUSCULO
func mult(x, y int) int {
	return x * y
}
