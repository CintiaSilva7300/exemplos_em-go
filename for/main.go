package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}
	fmt.Println("------------")

	numeros := []string{"um", "dois", "tres"}

	for indice, valor := range numeros {
		println(indice, valor)
	}

	fmt.Println("------------")
	i := 0
	for i < 10 {
		println(i)
		i++
	}

	fmt.Println("------------")
	// for {
	// 	println("loop infinito")
	// }
}