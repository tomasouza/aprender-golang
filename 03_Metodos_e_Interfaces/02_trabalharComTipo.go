package main

import (
	"fmt"
	"time"
)

// uma estrutura de empreendedor para utilizarmos como tipo
type Empreendedor struct {
	ID             int
	Nome           string
	Sobrenome      string
	DataNascimento time.Time
}

// um mapa de empreendedores que funcionará como nosso bando de dados em memoria
var Empreendedors = map[int]Empreendedor{}

// variavel utilitaria como ID da base de dados em memoria
var nextID = 0

// utilizando do nextID armazena as informacoes em um novo usuario na base de dados em memoria
func AdicionarEmpreendedor(nome, sobrenome string, dataNascimento time.Time) int {
	nextID++
	Empreendedors[nextID] = Empreendedor{
		ID:             nextID,
		Nome:           nome,
		Sobrenome:      sobrenome,
		DataNascimento: dataNascimento,
	}
	return nextID
}

// Busca um empreendedor pelo id em nossa base de dados em memoria
func BuscarEmpreendedor(id int) (Empreendedor, bool) {
	p, ok := Empreendedors[id]
	return p, ok
}

// Funcao que constroi um objeto de data a partir do dia, mes e ano.
func DMAToTime(dia int, mes time.Month, ano int) time.Time {
	return time.Date(ano, mes, dia, 0, 0, 0, 0, time.UTC)
}

func main() {
	e1ID := AdicionarEmpreendedor("Tomas", "Souza", DMAToTime(11, time.May, 1988))
	e2ID := AdicionarEmpreendedor("Anderson", "Silva", DMAToTime(11, time.May, 1988))
	e1, exists1 := BuscarEmpreendedor(e1ID)
	e2, exists2 := BuscarEmpreendedor(e2ID)
	fmt.Println(e1, exists1)
	fmt.Println(e2, exists2)
	e3, exists3 := BuscarEmpreendedor(2000)
	fmt.Println(e3, exists3)
}
