package main

import (
	"fmt"
)

var b float32 //global

type ID int

var ( //declaração em lote
	c int     = 40 // daclaração com inicialização
	d int          //valor padrão 0
	a bool         // false
	f float32      //+0.000000e+000
	s string       //"" espaço vazio

	i ID = 1
)

func main() {
	fmt.Printf("o tipo de %T", a)

}
