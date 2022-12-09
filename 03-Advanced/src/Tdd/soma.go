package main

import "fmt"

func main() {
	x := soma(1, 2, 3)
	fmt.Println(x)
}

func soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

// func mult() {
// 	x := soma(1, 2, 3)
// 	y := multiplic(10, 10)
// 	fmt.Println(x, y)
// }
