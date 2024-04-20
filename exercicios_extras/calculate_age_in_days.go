// Calculando a idade em dias.

//  Acabou não sendo tão facil mas deu certo kkk
package main

import "fmt"

func is_leap_year(actual_year int) bool {
	year_is_leap := false
	if actual_year%4 == 0 && actual_year%100 == 0 {
		year_is_leap = true
		fmt.Printf("O ano de %v é bissexto? %v", actual_year, year_is_leap)

	}
	return year_is_leap
}

func main() {
	actual_year := 2024
	age := 27
	birthday_year := actual_year - age
	var total_days int
	for x := 0; x < age; x++ {
		days := 0
		leap_year := is_leap_year(birthday_year + x)
		if leap_year {
			days = 366
		} else {
			days = 365
		}
		total_days += days
	}
	fmt.Println("Total de anos em dias: ", total_days)
}
