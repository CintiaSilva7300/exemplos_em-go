package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct { //site, coneverter json para struct(https://transform.tools/json-to-go)
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] { //os.Args[1:] pega todos os argumentos passados na execução(linha de comando)
		req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/") //fazer uma requisição HTTP

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
			
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
		}

		var data ViaCEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao decodificar resposta: %v\n", err)
		}
		
		fmt.Println(data) //ex de execução: go run main.go http://viacep.com.br/ws/13330-230/json/
	
		file, err := os.Create("cidade.txt") //cria o arquivo
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
		}

		defer file.Close() //fechar o arquivo
		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Logradouro: %s, Complemento: %s, Bairro: %s, Localidade: %s, UF: %s\n", data.Cep, data.Logradouro, data.Complemento, data.Bairro, data.Localidade, data.Uf)) //escreve no arquivo
		if err != nil { //verifica se deu tudo certo
			fmt.Fprintf(os.Stderr, "Erro ao escrever no arquivo: %\n", err)
		}

		fmt.Println("Arquivo criado com sucesso!")
		fmt.Println("Cidade: " + data.Localidade)
	}

}

//rodar: go run main.go 06550000