package main

import (
	"encoding/json"
	"io"
	"net/http"
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
	http.HandleFunc("/", BuscarCepHandelr) //definir a rota
	http.ListenAndServe(":8080", nil)//definir a porta
}

func BuscarCepHandelr(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Write([]byte("Hello World"))
	
	//exemplo de request: http://localhost:8080/?cep=01001000
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cep, error := BucaCep(cepParam)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(cep)//converte a struct para json

	//exemplo de response: {"cep":"01001-000","logradouro":"","complemento":"","bairro":"","localidade":"São Paulo","uf":"SP","ibge":"3550308","gia":"","ddd":"11","siafi":"7107"}
}

func BucaCep(cep string) (*ViaCEP, error) {
	resp, error := http.Get("http://viacep.com.br/ws/" + cep + "/json/")//fazer uma requisição HTTP para o CEP
	
	if error != nil {//verifica se deu tudo certo
		return nil, error
	}
	defer resp.Body.Close()

	body, error := io.ReadAll(resp.Body)//ler o corpo da resposta
	if error != nil {
		return nil, error
	}

	var c ViaCEP
	error = json.Unmarshal(body, &c)//converter json para struct
	if error != nil {
		return nil, error
	}
	return &c, nil
}

