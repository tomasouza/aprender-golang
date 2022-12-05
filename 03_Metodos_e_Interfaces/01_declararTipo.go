package main

import (
	"fmt"
	"time"
)

type Empregado struct {
	ID              int
	PrimeiroNome    string
	Sobrenome       string
	DataContratacao time.Time
}

func main() {
	//declarando um novo objeto Empregado
	e1 := Empregado{
		ID:              1,
		PrimeiroNome:    "Tomas",
		Sobrenome:       "Souza",
		DataContratacao: time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC),
	}
	e2 := Empregado{
		ID:              2,
		PrimeiroNome:    "Anderson",
		Sobrenome:       "Silva",
		DataContratacao: time.Date(2007, time.March, 30, 0, 0, 0, 0, time.UTC),
	}
	//Utilizando dos valores dos structs
	fmt.Println(e1.PrimeiroNome)
	fmt.Println(e2.DataContratacao)
	// Tomas se casou e mudou seu nome
	e1.Sobrenome = "Novas Brito"
	fmt.Println(e1.Sobrenome)
}
