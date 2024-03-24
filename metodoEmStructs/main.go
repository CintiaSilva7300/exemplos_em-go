package main

import "fmt"

type Endereco struct {
	Rua    string
	Bairro string
	Cidade string
	CEP    int
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c *Cliente) Desativar() {
	c.Ativo = false

	fmt.Println("O cliente",c.Nome,"foi desativado",c.Nome)
}

func main() {

	cintia := Cliente{
		Nome:  "Cintia",
		Idade: 27,
		Ativo: true,
	}

	cintia.Ativo = false
	cintia.Desativar()
}