package main

import (
	"fmt"
	"testing"
)

type test struct {
	data   []int
	answer int
}

func main() {
	values := soma(1, 2, 3, 5)
	fmt.Println(values)
}

func soma(i ...int) int {
	total := 0
	for _, v := range i {
		total = total + v
	}
	return total
}

// go test
func TestSoma(t *testing.T) {
	got := soma(1, 2, 3, 4)
	expected := 10
	if got != expected {
		t.Error("Expected:", expected, "Got:", got)
	}
}

// Esse vai jogar um erro com o t.Error()
//
//	func TestSomaWithError(t *testing.T) {
//		got := soma(1, 2, 3, 4, 5)
//		expected := 10
//		if got != expected {
//			t.Error("Expected:", expected, "Got:", got)
//		}
//	}
func TestSomaEmTabela(t *testing.T) {
	tests := []test{
		{data: []int{1, 2, 3}, answer: 6},
		{[]int{5, 5, 5}, 15},
		{[]int{3000, 2000, 1000}, 6000},
	}
	for _, v := range tests {
		got := soma(v.data...)
		expected := v.answer
		if got != expected {
			t.Error("Expected:", expected, "Got:", got)
		}
	}
}

// Um teste de exemplo
func ExampleSoma() {
	//  Isso aqui é um teste valido, ele precisa
	// retornar esses valores nos comentarios.
	fmt.Println(soma(1, 2, 3))
	fmt.Println(soma(3, 2, 1))
	fmt.Println(soma(2, 2, 3))
	// Output:
	// 6
	// 6
	// 7
}
