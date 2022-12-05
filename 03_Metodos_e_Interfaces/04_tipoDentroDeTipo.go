package main

import (
	"fmt"
	"time"
)

type Humano struct {
	Nome           string
	Sobrenome      string
	DataNascimento time.Time
}

type Mago struct {
	Humano
	Magias []string
}

func (mago *Mago) aprenderMagia(magia string) {
	mago.Magias = append(mago.Magias, magia)
}

func criarHumano(nome string, sobrenome string, dataNascimento time.Time) Humano {
	return Humano{Nome: nome, Sobrenome: sobrenome, DataNascimento: dataNascimento}
}

func evoluirEmMago(humano Humano) *Mago {
	return &Mago{Magias: []string{}, Humano: humano}
}

func novaData(dia int, mes time.Month, ano int) time.Time {
	return time.Date(ano, mes, dia, 0, 0, 0, 0, time.UTC)
}

func main() {
	humano := criarHumano("Tomas", "Souza", novaData(11, time.May, 1988))
	mago := evoluirEmMago(humano)
	mago.aprenderMagia("Haducken")
	fmt.Println("Humano:", humano)
	fmt.Println("Mago:", mago)
	fmt.Println("Nome do mago:", mago.Nome)
	fmt.Println("Sobrenome do mago:", mago.Sobrenome)
	fmt.Println("Data de nascimento do mago:", mago.DataNascimento)
	fmt.Println("Magias do mago:", mago.Magias)
}
