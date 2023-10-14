package main

import (
	"fmt"
)

type Endereco struct {
	logradouro string
	numero     int
	cidade     string
	estado     string
}

type Cliente struct {
	nome  string
	idade int
	ativo bool

	Endereco
}

// funçõesjhgjhgjhg
func main() {

	renato := Cliente{
		nome:  "Renato",
		idade: 40,
		ativo: true,
	}
	renato.cidade = "Garanhuns"
	renato.logradouro = "av. Orlando Wanderley"
	renato.numero = 80
	renato.estado = "PE"

	fmt.Printf(" Nome: %s\n Idade: %d\n Ativo:%t\nEndereco:\n Logradouro: %s\n N°: %d\n Cidade: %s\n Estado:%s\n", renato.nome, renato.idade, renato.ativo, renato.logradouro, renato.numero, renato.cidade, renato.estado)

}
