package main

import "fmt"

func media(lista []float64) float64 {
	total := 0.0
	for _, valor := range lista {
		total += valor
	}
	return total / float64(len(lista))
}

func main() {
	lista := []float64{98, 93, 77, 82, 83}
	fmt.Println(media(lista))
}

func uniqueFunc() {
	lista := []float64{98, 93, 77, 82, 83}
	total := 0.0
	for _, valor := range lista {
		fmt.Println(valor)
		total += valor
	}
	media := (total / float64(len(lista)))
	fmt.Println("Média da turma é: ", media)
}
