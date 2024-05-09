// - Utilize mutex para consertar a condição de corrida do exercício anterior.

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

	var mu sync.Mutex

	for i := 0; i < totaldegoroutines; i++ {
		go func() {
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Valor final:", contador)
}

//  verificar race conditions
//  go run -race exercicios-nv9/ex_04.go
