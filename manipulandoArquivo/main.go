package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo"))
	// tamanho, err := f.WriteString([]byte("Escrevendo dados no arquivo"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", tamanho)
	defer f.Close()
	

	//leitura de arquivo
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo)) //converte para string "slice de bytes"

	//leitura de pouco em pouco abrindo o aquivo
	arquivo2, err := os.Open("arquivo.txt") //abre o arquivo
	if err != nil {
		panic(err)
	}
	
	reader := bufio.NewReader(arquivo2) //cria(reader) o leitor
	buffer := make([]byte, 1)//cria o buffer de 10 bytes

	for {
		n, err := reader.Read(buffer)//le o buffer
		if err != nil {
			break
		}
		fmt.Println("Lido cada pedaço de 1 byte: ",string(buffer[:n])) //converte para string "slice de bytes" n é a posição do buffer
	}
}