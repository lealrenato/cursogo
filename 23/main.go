package main

import (
	"errors"
	"fmt"
)

// funçõesjhgjhgjhg
func main() {
	valor, err := sum(1, 50)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)
}

func sum(a int, b int) (int, error) {
	if a+b >= 50 {
		return 400, errors.New("a soma é maior que 50 ")
	}
	return a + b, nil
}
