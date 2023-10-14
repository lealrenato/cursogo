package main

import (
	"fmt"
)

type Cliente struct {
	nome  string
	idade int
	ativo bool
}

// funçõesjhgjhgjhg
func main() {

	renato := Cliente{
		nome:  "Renato",
		idade: 40,
		ativo: true,
	}

	fmt.Printf("Nome: %s\nIdade: %d\nAtivo:%t\n", renato.nome, renato.idade, renato.ativo)

}
