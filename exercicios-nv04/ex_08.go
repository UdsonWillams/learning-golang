// - Crie um map com key tipo string e value tipo []string.
//   - Key deve conter nomes no formato sobrenome_nome
//   - Value deve conter os hobbies favoritos da pessoa
//
// - Demonstre todos esses valores e seus indexes.
package main

import "fmt"

func main() {

	my_map := map[string][]string{
		"udson_willams": []string{"jogar ci es", "teste1"},
		"elida_maria":   []string{"viajar", "teste2"},
		"jose_eduardo":  []string{"vender produtos", "teste3"},
		"everton_lima":  []string{"andar de moto", "teste4"},
	}

	for k, v := range my_map {
		fmt.Printf("chave: %v \n", k)
		for i, hobbie := range v {
			fmt.Printf("index: %v e com valor: %v\n", i, hobbie)
		}
		fmt.Print("\n")
	}

}
