package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("http://www.google.com.br") //fazer uma chamada HTTP GET para o google

	if err != nil {
		panic(err)
	}

	defer req.Body.Close() //fechar o corpo da resposta, sempre no final do programa

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	// defer fmt.Println("Primeira linha") //usado para executar por ultimo
	// fmt.Println("Segunda linha")
	// fmt.Println("Terceira linha")
}