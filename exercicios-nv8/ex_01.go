// - Partindo do código abaixo,
// utilize marshal para transformar []user em JSON.
// - https://play.golang.org/p/U0jea43X55
// - Atenção! Tem pegadinha aqui.

// A pegadinha é que os valores precisam estar com a primeira letra
// EM MAIUSCULO para a lib do json fazer o marshal corretamente.
// Precisei deixar os valores nesse formato, User, First, Age, etc.
// CAMPOS COM LETRAS MINUSCULO NÃO SÃO EXPORTADOS PARA JSON.

package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	First string
	Age   int
}

func main() {
	u1 := User{
		First: "James",
		Age:   32,
	}

	u2 := User{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := User{
		First: "M",
		Age:   54,
	}

	users := []User{u1, u2, u3}

	fmt.Println(users)
	jason, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jason))
}
