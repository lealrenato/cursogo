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

type Pessoa interface {
	Desativar()
}

type Cliente struct {
	nome  string
	idade int
	ativo bool
	Endereco
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {

}

func (c Cliente) Desativar() {
	c.ativo = false
	fmt.Printf("cliente %s foi desativado \n", c.nome)
	fmt.Printf("%t", c.ativo)

}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
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

	fmt.Printf(" Nome: %s\n Idade: %d\n Ativo:%t\nEndereco:\n Logradouro: %s\n N°: %d\n Cidade: %s\n Estado:%s\n\n\n\n", renato.nome, renato.idade, renato.ativo, renato.logradouro, renato.numero, renato.cidade, renato.estado)

	//minhaempresa := Empresa{}
	Desativacao(renato)

	fmt.Printf("\n\n\n")

	fmt.Printf(" Nome: %s\n Idade: %d\n Ativo:%t\nEndereco:\n Logradouro: %s\n N°: %d\n Cidade: %s\n Estado:%s\n\n\n\n", renato.nome, renato.idade, renato.ativo, renato.logradouro, renato.numero, renato.cidade, renato.estado)

}
