// - Crie um struct "pessoa"
// - Crie uma função chamada mudeMe que tenha *pessoa como parâmetro.
// Essa função deve mudar um valor armazenado no endereço *pessoa.
//   - Dica: a maneira "correta" para fazer dereference de um valor
//     em um struct seria (*valor).campo
//   - Mas consta uma exceção na documentação.
//     Link: https://golang.org/ref/spec#Selectors
//   - "As an exception, if the type of x is a named pointer type and
//     (*x).f is a valid selector expression denoting a field (but not a method),
//   - → x.f is shorthand for (*x).f." ←
//   - Ou seja, podemos usar tanto o atalho p1.nome quanto o tradicional (*p1).nome
//
// - Crie um valor do tipo pessoa;
// - Use a função mudeMe e demonstre o resultado.
// Exercicio de ponteiro

package main

import "fmt"

type Pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

// Ideia que tive
func MudeMe_1(p Pessoa) Pessoa {
	pointer := &p.nome
	*pointer = "Um Novo Nome lhe foi dado"
	return p
}

// Formato passado pela professora.
func MudeMe_2(p *Pessoa) *Pessoa {
	(*p).nome = "Um Novo Nome lhe foi dado"
	return p
}

// Formato passado pela professora que a linguagem diz que também é aceitavel.
func MudeMe_3(p *Pessoa) *Pessoa {
	p.nome = "Um Novo Nome lhe foi dado"
	return p
}

func main() {

	pessoa1 := Pessoa{"Udson", "Willams", 27}
	pessoa2 := Pessoa{"João", "Maria", 27}
	pessoa3 := Pessoa{"Fernando", "Sousza", 27}

	fmt.Println(pessoa1)
	fmt.Println(pessoa2)
	fmt.Println(pessoa3)
	fmt.Println("Endereço das pessoas 1, 2 e 3")
	fmt.Println(&pessoa1.nome, &pessoa2.nome, &pessoa3.nome)

	fmt.Println("---------------------------------------------")

	mudeme_1 := MudeMe_1(pessoa1)
	mudeme_2 := MudeMe_2(&pessoa2)
	mudeme_3 := MudeMe_3(&pessoa3)

	fmt.Println(mudeme_1)  // DO tipo pessoa
	fmt.Println(*mudeme_2) // Tipo ponteiro de pessoa
	fmt.Println(*mudeme_3) // Tipo ponteiro de pessoa
	fmt.Println("Endereço de pessoas 1, 2 e 3")
	fmt.Println(&pessoa1.nome, &pessoa2.nome, &pessoa3.nome)
}
