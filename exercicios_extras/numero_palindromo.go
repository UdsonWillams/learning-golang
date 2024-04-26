package main

import "fmt"

func main() {

	fmt.Println(NumeroPalindromo(121))
	fmt.Println(NumeroPalindromo(222))
	fmt.Println(NumeroPalindromo(313))
	fmt.Println(NumeroPalindromo(456))
	fmt.Println(NumeroPalindromo(54545))
	fmt.Println(NumeroPalindromo(634634))
	fmt.Println(NumeroPalindromo(77777))
	fmt.Println(NumeroPalindromo(8))
	fmt.Println(NumeroPalindromo(9))
	fmt.Println(NumeroPalindromo(1000021))
	fmt.Println(NumeroPalindromo(0))
}

func NumeroPalindromo(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 || x%10 == 0 {
		return false
	}

	reversed := 0
	for x > reversed {
		reversed = reversed*10 + x%10
		x /= 10
	}
	return (x == reversed) || (x == reversed/10)
}
