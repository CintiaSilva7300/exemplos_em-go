package main

import "fmt"

type Endereco struct {
	Rua    string
	Bairro string
	Cidade string
	CEP    int
}

type Pessoa interface {
	Desativar()
}

type Empresa struct {
	nome string
}

// Desativar implements Pessoa.
func (e Empresa) Desativar() {
	panic("unimplemented")
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c *Cliente) Desativar() {
	// c.Ativo = false

	// fmt.Println("O cliente",c.Nome,"foi desativado",c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {

	cintia := Cliente{
		Nome:  "Cintia",
		Idade: 27,
		Ativo: true,
	}

	minhaEmpresa := Empresa{}
	Desativacao(minhaEmpresa)

	Desativacao(&cintia)

	fmt.Println("Nome:", cintia.Nome, ",", "Idade:", cintia.Idade, ",", "Ativo:", cintia.Ativo)

}
