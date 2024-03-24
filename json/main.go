package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero float64 `json:"numero"`
	Saldo  float64 `json:"saldo"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}

	res, err := json.Marshal(conta) //Marshal transforma o objeto em JSON
	
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		fmt.Println(err)
	}

	jsonPuro := []byte(`{"numero":1, "saldo":2}`)
	var contaX Conta
	err =  json.Unmarshal(jsonPuro, &contaX)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(contaX)
}