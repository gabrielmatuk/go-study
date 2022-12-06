package main

import "fmt"

func main() {
	x := 5
	if x == 2 || x == 3 {
		fmt.Println("Sim, X é 2 ou 3")
	} else {
		fmt.Println("Nenhuma das opções")
	}
}
