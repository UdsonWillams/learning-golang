package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	senha := "Joãozin123"
	senha_errada := "Joãozin1233"

	sb, err := bcrypt.GenerateFromPassword([]byte(senha), 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(sb))

	// Senha correta retorna nda, nil
	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senha)))
	// Senha incorreta, retorna um error
	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senha_errada)))
}
