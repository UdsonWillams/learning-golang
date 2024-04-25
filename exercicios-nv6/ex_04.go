// - Crie um tipo struct "pessoa" que contenha os campos:
//     - nome
//     - sobrenome
//     - idade
// - Crie um método para "pessoa" que demonstre o nome completo e a idade;
// - Crie um valor de tipo "pessoa";
// - Utilize o método criado para demonstrar esse valor.

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa) ShowPeople() {
	fmt.Printf("Nome: %v %v e idade: %v", p.nome, p.sobrenome, p.idade)
}

func main() {
	pessoa := pessoa{"Udson", "Willams", 27}
	pessoa.ShowPeople()
	fmt.Println()
}
