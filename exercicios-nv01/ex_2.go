// Use var para declarar três variáveis. Elas devem ter
// package-level scope. Não atribua valores a estas variáveis.
// Utilize os seguintes identificadores e tipos para estas variáveis:
// 1. Id "x" deverá ter tipo int
// 2. Id "y" deverá ter tipo string
// 3. Id "z" deverá ter tipo bool

// na função main:
// 1. Demonstre os valores de cada identificador.
// 2. O compilador atribuiu valores para essas variáveis. Como esses valores se chamam?
package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Println(x, ">", y, "<", z)
	fmt.Printf("valor de x: %v |valor de y: %v |valor de z: %v |\n", x, y, z)
	fmt.Println("Os valores atribuidos a essas variaveis foram o valor 0 especifico de cada tipo")
}

