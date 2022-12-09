package main

import "testing"

// go test
func TestSoma(t *testing.T) {
	test := soma(3, 2, 1)
	resultado := 6
	if test != resultado {
		t.error("Valor esperado: ", resultado, "valor retornado", test)
	}
}
