package main

import "fmt"

const ebulicaoF = 212.0
const ebulicaoK = 373.2

func main() {
	var tempF = ebulicaoF
	var tempC = (tempF - 32) * 5 / 9
	tempK := ebulicaoK
	tempCK := (ebulicaoK - 273)
	fmt.Println("A temperatura de ebulição da água em F = ", tempF)
	fmt.Println("A temperatura da ebulição da água em C = ", tempC)
	fmt.Printf("A temperatura de ebulição da água em K é: %g e o tempo de ebulição da água em C = %g", tempK, tempCK)

}
