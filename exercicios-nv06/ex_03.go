// - Utilize a declaração defer de maneira que demonstre que sua execução só ocorre ao final do contexto ao qual ela pertence.

package main

import "fmt"

func main() {
	fmt.Println("1° valor")
	fmt.Println("2° valor")
	defer fmt.Println("- 3° na ordem, mas ultimo valor por conta do defer")
	fmt.Println("3° valor")
	fmt.Println("4° valor")
}
