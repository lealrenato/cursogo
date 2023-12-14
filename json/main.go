package main

import (
	"encoding/json"
	"os"
)

// Json

type Conta struct {
	Numero int `json:"numero"` //tag
	Saldo  int `json:"saldo"`
}

func main() {
	//tranformar struct em json guardando em uma variavel
	conta := Conta{Numero: 1, Saldo: 100}

	res, err := json.Marshal(conta)

	if err != nil {
		panic(err)
	}

	println(string(res))

	//tranformar struct em json passando valor sem guardar em variavel passando para um arquivo ou terminal(encoder)
	err = json.NewEncoder(os.Stdout).Encode(conta) //retorna um erro
	if err != nil {
		panic(err)
	}

	//decodificando json para struct
	jsonPuro := []byte(`{"n":2,"s":200}`)

	var contaX Conta

	err = json.Unmarshal(jsonPuro, &contaX) //retorna um erro

	if err != nil {
		panic(err)
	}

	println(contaX.Saldo)
}
