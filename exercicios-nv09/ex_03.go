// - Utilizando goroutines, crie um programa incrementador:
//   - Tenha uma variável com o valor da contagem
//   - Crie um monte de goroutines, onde cada uma deve:
//   - Ler o valor do contador
//   - Salvar este valor
//   - Fazer yield da thread com runtime.Gosched()
//   - Incrementar o valor salvo
//   - Copiar o novo valor para a variável inicial
//   - Utilize WaitGroups para que seu programa não finalize antes de suas goroutines.
//   - Demonstre que há uma condição de corrida utilizando a flag -race
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	contador := 0
	totaldegoroutines := 30

	var wg sync.WaitGroup
	wg.Add(totaldegoroutines)

	for i := 0; i < totaldegoroutines; i++ {
		go func() {
			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Valor final:", contador)
}

//  verificar race conditions
//  go run -race exercicios-nv9/ex_03.go
