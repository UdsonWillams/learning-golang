package main

import (
	"fmt"
	"testing"
)

type test1 struct {
	data   []int
	answer int
}

func main() {
	values := soma(1, 2, 3, 5)
	fmt.Println(values)
}

func soma1(i ...int) int {
	total := 0
	for _, v := range i {
		total = total + v
	}
	return total
}

func multiplicacao(i ...int) int {
	total := 1
	for _, v := range i {
		total = total * v
	}
	return total
}

func divisao(i ...int) int {
	total := 1
	for _, v := range i {
		total = total / v
	}
	return total
}

func subtracao(i ...int) int {
	total := 0
	for _, v := range i {
		total = total - v
	}
	return total
}

func BenchmarkSoma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		soma1(100, 100)
	}
}

func BenchmarkMultiplicacao(b *testing.B) {
	for i := 0; i < b.N; i++ {
		multiplicacao(100, 100)
	}
}

func BenchmarkDivisao(b *testing.B) {
	for i := 0; i < b.N; i++ {
		divisao(100, 100)
	}
}

func BenchmarkSubtracao(b *testing.B) {
	for i := 0; i < b.N; i++ {
		subtracao(100, 100)
	}
}
