package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func main() {

	fmt.Println(pessoa{"Gabriel", 54})
	fmt.Println(pessoa{"Marcos", 21})

}
