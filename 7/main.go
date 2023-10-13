package main

import (
	"fmt"
)

func main() {

	salarios := map[string]int{"Renato": 3000, "João": 2300, "Maria": 2900}
	//sal := make(map[string]int)
	//sal1 := map[string]int{}

	fmt.Println(salarios["Renato"])
	fmt.Println(salarios)

	delete(salarios, "Renato")
	fmt.Println(salarios)

	salarios["Renato"] = 5000
	fmt.Println(salarios)

	salarios["Renato"] = 6000
	fmt.Println(salarios)

	for nome, salario := range salarios {

		fmt.Printf("O salario de %s é R$%d\n", nome, salario)
	}

	for _, salario := range salarios {

		fmt.Printf("O salario  é R$%d\n", salario)
	}

}
