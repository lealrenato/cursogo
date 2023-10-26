package main

// funçõesjhgjhgjhg
func main() {

	a := 10
	println("valor de", a)
	println("ponteiro de a", &a)

	var b *int = &a

	println("valor de b que recebe endereco de a", b)

	println("valor contido no enderereço que b recebe", *b)

	*b = 30

	println("valor de a auterado apartir de *b ", a)

}
