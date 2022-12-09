package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Nome    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

type Pokemon struct {
	Numero  int            `json:"entry_numbers"`
	Especie PokemonSpecies `json:"pokemon_species"`
}

type PokemonSpecies struct {
	Nome string `json:"name"`
}

const POKE_URL string = "http://pokeapi.co/api/v2/pokedex/kanto/"

func main() {
	response, err := http.Get(POKE_URL)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	fmt.Println(responseObject.Nome)
	fmt.Println(responseObject.Pokemon)

	for _, Pokemon := range responseObject.Pokemon {
		fmt.Println(Pokemon.Especie.Nome)
	}
}
