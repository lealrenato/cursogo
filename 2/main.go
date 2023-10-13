package main

var b float32 //global

var ( //declaração em lote
	c int     = 40 // daclaração com inicialização
	d int          //valor padrão 0
	a bool         // false
	f float32      //+0.000000e+000
	s string       //"" espaço vazio
)

func main() {
	var b float32 //escopo local

	println(b)
}
