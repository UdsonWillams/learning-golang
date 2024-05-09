// - Utilizando este código: https://play.golang.org/p/3W69TH4nON
// - ...remova o underscore e verifique e lide com o erro de maneira apropriada.
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, erro := json.Marshal(p1)
	fmt.Println(erro)
	if erro != nil {
		log.Println("Aconteceu um erro aqui amigo!")
	}
	fmt.Println(string(bs))

}
