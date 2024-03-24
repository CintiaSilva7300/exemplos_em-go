package main

import (
	"fmt"
)

type Endereco struct {
	Rua string
	Bairro string
	Cidade string
	CEP int
}

type Cliente struct {
	Nome string
	Idade int
	Ativo bool
	Address Endereco
}

func main() {
	
	cintia := Cliente {
		Nome: "Cintia",
		Idade: 27,
		Ativo: true,
	}

	cintia.Ativo = false
	cintia.Address.Cidade = "Sao Paulo"
	cintia.Address.Bairro = "Centro"

	fmt.Println("Nome:",cintia.Nome,",", "Idade:",cintia.Idade,",", "Ativo:", cintia.Ativo)
	fmt.Println("Cidade:",cintia.Address.Cidade,",", "Bairro:",cintia.Address.Bairro)
}