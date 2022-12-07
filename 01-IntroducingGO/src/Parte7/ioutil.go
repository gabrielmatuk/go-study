package main

import (
	"io/ioutil"
	"log"
)

func main() {
	message := []byte("Olá mundo")
	err := ioutil.WriteFile("arquivo", message, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
