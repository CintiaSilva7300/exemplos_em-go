package main

import "fmt"

type Person struct {
	name string
	age int
}

func main() {
	var meuArray [3] int

	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30
	fmt.Println(meuArray[0])
	fmt.Println(meuArray[1])
	fmt.Println(meuArray[2])
	fmt.Println("------------")

	fmt.Println("ultima posição",len(meuArray)-1)
	fmt.Println("tamanho",len(meuArray))
	fmt.Println("ultimo posição",meuArray[len(meuArray)-1])
	fmt.Println("------------")

	//percorrer array
	for i := 0; i < len(meuArray); i++ {
		fmt.Println("percorrendo array",meuArray[i])
	}

	fmt.Println("------------")

	//percorrer array
	for i, v := range meuArray { 
		fmt.Println("percorrendo array",i,v)
	}
}