package main

import (
	"fmt"
	"os"
)

var sequence = []int{12, 13, 14, 15, 16, 17, 18, 19, 22, 23, 24, 25, 26, 27, 28, 29, 32, 33, 34, 35, 36, 37, 38, 39, 42, 43, 44, 45, 46, 47, 48, 49}

var list1 = []int{12, 13, 14, 15, 22, 23, 24, 25}
var list2 = []int{16, 17, 18, 19, 26, 27, 28, 29}
var list3 = []int{32, 33, 34, 35, 42, 43, 44, 45}
var list4 = []int{36, 37, 38, 39, 46, 47, 48, 49}

func isNeighborInRow(a, b int) bool {
	return (a+1)%8 != 0 && a+1 == b
}

func isEven(num int) bool {
	return num%2 == 0
}

func isValidCombination(comb []int) bool {
	evenCount, oddCount := 0, 0

	for i, num1 := range comb {
		for _, num2 := range comb[i+1:] {
			if isNeighborInRow(num1, num2) || num1+1 == num2 {
				return false
			}
		}

		if isEven(num1) {
			evenCount++
		} else {
			oddCount++
		}

		if evenCount > 3 || oddCount > 3 {
			return false
		}
	}

	return evenCount == 3 && oddCount == 3
}

func countOccurrencesInList(list []int, comb []int) int {
	count := 0
	for _, num := range comb {
		if containsNumber(list, num) {
			count++
		}
	}
	return count
}

func isValidWithLists(comb []int) bool {
	countList1 := countOccurrencesInList(list1, comb)
	countList2 := countOccurrencesInList(list2, comb)
	countList3 := countOccurrencesInList(list3, comb)
	countList4 := countOccurrencesInList(list4, comb)

	return (countList1 >= 1 && countList1 <= 2) &&
		(countList2 >= 1 && countList2 <= 2) &&
		(countList3 >= 1 && countList3 <= 2) &&
		(countList4 >= 1 && countList4 <= 2)
}

func containsNumber(list []int, num int) bool {
	for _, n := range list {
		if n == num {
			return true
		}
	}
	return false
}

func generateCombinations() [][]int {
	var combinations [][]int

	backtrack(sequence, 0, []int{}, &combinations)

	return combinations
}

func backtrack(remaining []int, index int, current []int, result *[][]int) {
	if len(current) == 6 {

		if isValidCombination(current) && isValidWithLists(current) {
			*result = append(*result, append([]int{}, current...))
		}
		return
	}

	for i := index; i < len(remaining); i++ {
		// Adiciona o elemento atual à combinação
		current = append(current, remaining[i])

		// Chama recursivamente para o restante da combinação
		backtrack(remaining, i+1, current, result)

		// Remove o último elemento para tentar diferentes combinações
		current = current[:len(current)-1]
	}
}

func writeToFile(combinations [][]int) error {
	file, err := os.Create("combinacoes.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	for i, combination := range combinations {
		fmt.Printf("Combinação %d: %v\n", i+1, combination)
		_, err := fmt.Fprintln(file, combination)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	combinations := generateCombinations()

	err := writeToFile(combinations)
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}

	fmt.Println("Combinações foram salvas em 'combinacoes.txt'")
}
