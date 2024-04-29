package main

import (
	"fmt"
	"sort"
)

type carro struct {
	nome     string
	potencia int
	consumo  int
}

//  criamos um tipo carro

type ordenarPorPotencia []carro

//  depois um tipo criamos um tipo slice de carro

func (x ordenarPorPotencia) Len() int           { return len(x) }
func (x ordenarPorPotencia) Less(i, j int) bool { return x[i].potencia < x[j].potencia }
func (a ordenarPorPotencia) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// criamos os metodos ligados a esse slice, e o metodo Sort
// vai utiliza-los para fazer as ordenações.

// Len == tamanho do slice, less == menor ou maior, swap == mudar os valores de lugar.
// Isso é repetido abaixo e é assim que criamos tipos
// que podem ser ordenados com o Sort.

// Lembrando que o pacote Sort possui um Sort.Slice()
// Pra uma simples slice, já é o suficiente.

type ordenarPorConsumo []carro

func (x ordenarPorConsumo) Len() int           { return len(x) }
func (x ordenarPorConsumo) Less(i, j int) bool { return x[i].consumo > x[j].consumo }
func (a ordenarPorConsumo) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ordenarPorLucroProDonoDoPosto []carro

func (x ordenarPorLucroProDonoDoPosto) Len() int           { return len(x) }
func (x ordenarPorLucroProDonoDoPosto) Less(i, j int) bool { return x[i].consumo < x[j].consumo }
func (a ordenarPorLucroProDonoDoPosto) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {

	carros := []carro{carro{"Chevete", 50, 8},
		carro{"Porsche", 300, 5},
		carro{"Fusca", 20, 30},
	}

	fmt.Println("Inicial:\n", carros)

	sort.Sort(ordenarPorPotencia(carros))

	fmt.Println("Potência:\n", carros)

	sort.Sort(ordenarPorConsumo(carros))

	fmt.Println("Consumo:\n", carros)

	sort.Sort(ordenarPorLucroProDonoDoPosto(carros))

	fmt.Println("Lucro:\n", carros)

}
