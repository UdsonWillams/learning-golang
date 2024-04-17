// - Crie um loop utilizando a sintaxe: for condition {}
// - Utilize-o para demonstrar os anos desde que você nasceu.

package main

import "fmt"

func main() {

	my_year := 1996
	actual_year := 2024
	for {
		fmt.Println(my_year)
		my_year++
		if my_year > actual_year {
			break
		}
	}
}
