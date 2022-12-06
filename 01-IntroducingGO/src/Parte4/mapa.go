package main

import "fmt"

func main() {

	x := make(map[string]int)
	x["Chave"] = 10

	y := make(map[int]int)
	y[1] = 20

	elemento := make(map[string]string)
	elemento["H"] = "Hidrogênio"
	elemento["He"] = "Hélio"
	elemento["Li"] = "Lítio"
	fmt.Println(x["Chave"])
	fmt.Println(y[1])
	fmt.Println(elemento["Li"])
}
