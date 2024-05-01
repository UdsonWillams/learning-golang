// - Esse exercício vai reforçar seus conhecimentos sobre conjuntos de métodos.
//   - Crie um tipo para um struct chamado "pessoa"
//   - Crie um método "falar" para este tipo que tenha um receiver ponteiro (*pessoa)
//   - Crie uma interface, "humanos", que seja implementada por tipos com o método "falar"
//   - Crie uma função "dizerAlgumaCoisa" cujo parâmetro seja do tipo "humanos" e que chame o método "falar"
//   - Demonstre no seu código:
//   - Que você pode utilizar um valor do tipo "*pessoa" na função "dizerAlgumaCoisa"
//   - Que você não pode utilizar um valor do tipo "pessoa" na função "dizerAlgumaCoisa"
package main

import "fmt"

type pessoa struct {
	Name string
}

func (p *pessoa) falar() {
	fmt.Println("Olá, sou o " + p.Name)
}

type humanos interface {
	falar()
}

func dizerAlgumaCoisa(h humanos) {
	h.falar()
}

func main() {
	p1 := pessoa{Name: "Udson Willams"}
	// tbm pode ser dessa forma
	p1.falar()
	(&p1).falar()

	fmt.Println(`Posso usar o tipo "*pessoa" na func dizerALgumaCOisa`)
	dizerAlgumaCoisa(&p1)
	fmt.Println("Não posso usar o tipo pessoa na func dizerALgumaCOisa")
	// dizerAlgumaCoisa(p1)

}

//   - Que você pode utilizar um valor do tipo "*pessoa" na função "dizerAlgumaCoisa"
//   - Que você não pode utilizar um valor do tipo "pessoa" na função "dizerAlgumaCoisa"
