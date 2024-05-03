package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {

	roman_values := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	roman_values_precedents := map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	unit_values := len(s)
	var total int
	var next_roman_value string

	if exist, ok := roman_values[s]; ok {
		return exist
	} else {
		for i := 0; i < unit_values; i++ {
			roman_value := string(s[i])
			if i+1 == unit_values {
				next_roman_value = string(s[i])
			} else {
				next_roman_value = string(s[i+1])
			}
			next := roman_value + next_roman_value
			if roman_value, ok := roman_values_precedents[next]; ok {
				total += roman_value
				i++
				continue
			}
			value := roman_values[roman_value]
			total += value
		}
		return total
	}
}
