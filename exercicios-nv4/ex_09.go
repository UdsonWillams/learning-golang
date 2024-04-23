// - Utilizando o exercício anterior, adicione uma entrada ao map e demonstre o map inteiro utilizando range.

package main

import "fmt"

func main() {

	my_map := map[string][]string{
		"udson_willams": []string{"jogar ci es", "teste1"},
		"elida_maria":   []string{"viajar", "teste2"},
		"jose_eduardo":  []string{"vender produtos", "teste3"},
		"everton_lima":  []string{"andar de moto", "teste4"},
	}

	my_map["marcos_souza"] = []string{"empinar de moto", "teste5"}

	for k, v := range my_map {
		fmt.Printf("chave: %v \n", k)
		for i, hobbie := range v {
			fmt.Printf("index: %v e com valor: %v\n", i, hobbie)
		}
		fmt.Print("\n")
	}

}
