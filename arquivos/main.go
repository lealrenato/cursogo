package main

import (
	"bufio"
	"fmt"
	"os"
)

// Manipulação de arquivos
func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}
	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo"))
	//tamanho, err := f.WriteString("ola mundo")
	if err != nil {
		panic(err)
	}
	fmt.Printf("arquivo criado com sucesso tamanho %d \n", tamanho)
	f.Close()

	///Leitura
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	//Leitura de pouco em pouco
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 3)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}
	arquivo2.Close()
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
