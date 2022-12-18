package main

import "fmt"

type primeiraPessoa struct {
	nome  string
	idade int
}

// Golang permite você atribuir uma estrutura a outra e comparar entre elas
// Desde que seus atributos possuam os mesmos nomes, sigam a mesma ordem e os mesmos tipos
func main() {
	p := primeiraPessoa{
		nome:  "Tomas Souza",
		idade: 34,
	}

	var s struct {
		nome  string
		idade int
	}
	//Atribuindo a estrutura primeiraPessoa a outra estrutura de nome s
	s = p
	fmt.Println("Segunda Pessoa:", s)
	//O Compilador entenderá como se os dois fossem iguais
	fmt.Println("As duas pessoas são iguais:", p == s)
	s.nome = "Milena"
	fmt.Println("As duas pessoas ainda são iguais:", p == s)
}
