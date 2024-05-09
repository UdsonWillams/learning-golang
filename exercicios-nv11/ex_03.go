// - Crie um struct "erroEspecial" que implemente a interface builtin.error.
// - Crie uma função que tenha um valor do tipo error como parâmetro.
// - Crie um valor do tipo "erroEspecial" e passe-o para a função da instrução anterior.
package main

import (
	"fmt"
)

type erroEspecial struct {
	msg   string
	extra string
}

func (err erroEspecial) Error() string {
	return fmt.Sprintf("Algo inesperado aconteceu: %v | %v", err.msg, err.extra)
}

func errorType(err error) {
	fmt.Println("Funcão Error Type")
	fmt.Println(err.Error())
}

func main() {

	meuErro := erroEspecial{msg: "Olá, sou um erro criado!", extra: "Isso não deveria acontecer"}
	errorType(meuErro)
}
