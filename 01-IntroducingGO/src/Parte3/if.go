package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("O número:", i, "É PAR")
		} else {
			fmt.Println("O número", i, "É ÍMPAR")
		}
		fmt.Println(i)
	}
}
