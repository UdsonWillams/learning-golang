// - Crie um novo tipo: veículo
//   - O tipo subjacente deve ser struct
//   - Deve conter os campos: portas, cor
//
// - Crie dois novos tipos: caminhonete e sedan
//   - Os tipos subjacentes devem ser struct
//   - Ambos devem conter "veículo" como struct embutido
//   - O tipo caminhonete deve conter um campo bool chamado "traçãoNasQuatro"
//   - O tipo sedan deve conter um campo bool chamado "modeloLuxo"
//
// - Usando os structs veículo, caminhonete e sedan:
//   - Usando composite literal, crie um valor de tipo caminhonete e dê valores a seus campos
//   - Usando composite literal, crie um valor de tipo sedan e dê valores a seus campos
//
// - Demonstre estes valores.
// - Demonstre um único campo de cada um dos dois.
package main

import "fmt"

type Veiculo struct {
	portas string
	cor    string
}

type Caminhonete struct {
	Veiculo
	traçãoNasQuatro bool
}
type Sedan struct {
	Veiculo
	modeloLuxo bool
}

func main() {

	caminhonete := Caminhonete{Veiculo{portas: "2 portas", cor: "preto"}, true}
	caminhonete_composta := Caminhonete{
		Veiculo:         Veiculo{portas: "4 portas", cor: "azul"},
		traçãoNasQuatro: false,
	}

	sedan := Sedan{Veiculo{portas: "2 portas", cor: "branco"}, false}
	sedan_composto := Sedan{
		Veiculo:    Veiculo{portas: "4 portas", cor: "azul"},
		modeloLuxo: true,
	}

	fmt.Println(caminhonete, sedan)
	fmt.Println(caminhonete_composta, sedan_composto)

	fmt.Printf("\n Cor da caminhonete: %v | tem tração? %v", caminhonete.cor, caminhonete.traçãoNasQuatro)
	fmt.Printf("\n Cor do Sedan: %v | modelo de luxo? %v", sedan.cor, sedan.modeloLuxo)
	fmt.Println()
}
