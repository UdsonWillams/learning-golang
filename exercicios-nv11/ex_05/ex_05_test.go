// - Nos capítulos seguintes, uma das coisas que veremos é testes.
// - Para testar sua habilidade de se virar por conta própria... desafio:
//   - Utilizando as seguintes fontes: https://godoc.org/testing & http://www.golang-book.com/books/intr...
//   - Tente descobrir por conta própria como funcionam testes em Go.
//   - Pode usar tradutor automático, pode rodar código na sua máquina, pode procurar no Google. Vale tudo.
//   - O exercício é: crie um teste simples de uma função ou método ou pedaço qualquer de código.
package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var v float64

type sqrtError1 struct {
	lat  string
	long string
	err  error
}

func (se sqrtError1) Error() string {
	return fmt.Sprintf("math error: %v - %v - %v", se.lat, se.long, se.err)
}

func funcToTest(f float64) (float64, error) {
	if f < 0 {
		return f, sqrtError1{
			lat:  "10",
			long: "21",
			err:  errors.New("Valor passado menor que 0"),
		}
	}
	return 42, nil
}

func TestSqrtError(t *testing.T) {
	v := -10.0
	_, err := funcToTest(v)
	assert.Equal(t, err.Error(), "math error: 10 - 21 - Valor passado menor que 0")
}

func TestSqrtSuccess(t *testing.T) {
	v := 42.0
	response, _ := funcToTest(v)
	assert.Equal(t, v, response)
}
