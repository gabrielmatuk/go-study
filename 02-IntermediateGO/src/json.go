package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Social struct {
	Midias string `json:"midias"`
}

type Usuario struct {
	Nome    string `json:"nome"`
	Idade   int    `json:"idade"`
	Salario int    `json:"salario"`
	Social  Social `json:"Social"`
}

type Usuarios struct {
	Usuarios []Usuario `json:"individuos"`
}

func main() {
	jsonFile, err := os.Open("usuarios.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Arquivo aberto com sucesso")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var usuarios Usuarios
	json.Unmarshal(byteValue, &usuarios)

	for i := 0; i < len(usuarios.Usuarios); i++ {
		fmt.Println("Usuário Nome: " + usuarios.Usuarios[i].Nome)
		fmt.Println("Usuário idade: " + strconv.Itoa(usuarios.Usuarios[i].Idade))
		fmt.Println("Usuário salario: " + strconv.Itoa(usuarios.Usuarios[i].Salario))

	}
}
