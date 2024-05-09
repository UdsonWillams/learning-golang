// - Crie um programa que demonstra seu OS e ARCH.
// - Rode-o com os seguintes comandos:
//   - go run == roda o arquivo passado, com ajuda do compilador
//   - go build == compila o programa, sendo possivel executalo no linux, windows, mac, etc.
//     depois de buildado e compilado usar o comando de execução do sistema, no linux: ./ex_06
//   - go install == instala o arquivo compilado no bin do diretorio gopath, sendo possivel
//     chamar o valor como um programa, por ex, agora passamos apenas: ex_06 pro terminal.
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Seu OS é:\t", runtime.GOOS)
	fmt.Println("Seu ARCH é:\t", runtime.GOARCH)
}

//   - go run == roda o arquivo passado, com ajuda do compilador
//   - go build == compila o programa, sendo possivel executalo no linux, windows, mac, etc.
//     depois de buildado e compilado usar o comando de execução do sistema, no linux: ./ex_06
//   - go install == instala o arquivo compilado no bin do diretorio gopath, sendo possivel
//     chamar o valor como um programa, por ex, agora passamos apenas: ex_06 pro terminal.
