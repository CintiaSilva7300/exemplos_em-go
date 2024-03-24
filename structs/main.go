package main

import (
	"fmt"
)

type Cliente struct {
	Nome string
	Idade int
	Ativo bool
}

func main() {
	
	cintia := Cliente {
		Nome: "Cintia",
		Idade: 27,
		Ativo: true,
	}

	cintia.Ativo = false

	fmt.Println("Nome:",cintia.Nome,",", "Idade:",cintia.Idade,",", "Ativo:", cintia.Ativo)
}