// - Alem da goroutine principal, crie duas outras goroutines.
// - Cada goroutine adicional devem fazer um print separado.
// - Utilize waitgroups para fazer com que suas goroutines finalizem antes de o programa terminar

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	number_of_routines := 3
	wg.Add(number_of_routines)
	go GoRoutine1()
	go GoRoutine2()
	go GoRoutine3()
	wg.Wait()
}

func GoRoutine1() {
	fmt.Println("go routine 1")
	wg.Done()
}

func GoRoutine2() {
	fmt.Println("go routine 2")
	wg.Done()
}

func GoRoutine3() {
	fmt.Println("go routine 3")
	wg.Done()
}

// Exemplo da professora
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var wg sync.WaitGroup

// func main() {
// 	novasgoroutines(100)
// 	wg.Wait()
// }

// func novasgoroutines(i int) {
// 	wg.Add(i)
// 	for j := 0; j < i; j++ {
// 		x := j
// 		go func(i int) {
// 			fmt.Println("Eu sou a goroutine número:", i+1)
// 			wg.Done()
// 		}(x)
// 	}
// }
