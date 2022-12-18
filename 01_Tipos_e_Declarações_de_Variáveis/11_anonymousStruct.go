package main

import "fmt"

var pessoa struct {
	nome            string
	idade           int
	animalEstimacao string
}

func main() {
	pessoa.nome = "Tomas Souza"
	pessoa.idade = 34
	pessoa.animalEstimacao = "Gato"

	fmt.Println("Pessoa:", pessoa)

	animal := struct {
		nome    string
		especie string
	}{
		nome:    "Eufelia",
		especie: "Cachorro",
	}

	fmt.Println("Animal:", animal)
}
