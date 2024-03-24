package main

import "fmt"

type MyNumeber int

type Numero interface {
	~int | ~float64 
}

func soma[T Numero ](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println(soma(m))

	m2 := map[string]float64{
		"a": 1.1,
		"b": 2.2,
		"c": 3.3,
	}

	fmt.Println(soma(m2))

	m3 := map[string]MyNumeber{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	fmt.Println(soma(m3))

	// fmt.Println(soma(m3))
	println(Compara(10, 10))
}
