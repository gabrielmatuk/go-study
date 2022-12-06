package main

import "fmt"

func main() {
	arr := [7]float64{1, 2, 3, 4, 5, 6, 7}
	fatiaNova := []int{1, 2, 3}
	fatia := make([]float64, 3, 5)

	fatia2 := arr[2:5]
	fatia3 := append(fatiaNova, 4, 5, 6)

	fatiaCopia := make([]int, 2)
	copy(fatiaCopia, fatiaNova)

	fmt.Println(fatia)
	fmt.Println(fatia2)
	fmt.Println(fatiaNova)
	fmt.Println(fatia3)

	fmt.Println(fatiaCopia)

}

//acrescentar elemento
