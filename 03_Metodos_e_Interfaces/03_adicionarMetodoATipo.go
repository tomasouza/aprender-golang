package main

import (
	"fmt"
	"time"
)

// Monta uma estrutura de ajudante para trabalharmos como nossa entidade
type Ajudante struct {
	ID          int
	Nome        string
	Sobrenome   string
	DataChegada time.Time
}

// Cria um tipo AjudantesDataAdapter para que utilizemos para construir os metodos
// Estes metodos serao adicionados ao tipo AjudantesDataAdapter
type AjudantesDataAdapter struct {
	ajudantes map[int]Ajudante
	//more2
	nextID int
}

// Funcao para criar um novo AjudantesDataAdapter
func NovoAjudantesDataAdapter() *AjudantesDataAdapter {
	return &AjudantesDataAdapter{
		ajudantes: map[int]Ajudante{},
		//more3
		nextID: 0,
	}
}

// Metodo que adiciona um Ajudante a estrutura de dados do AjudantesDataAdapter
// Retorna um ID deste novo Ajudante
func (adapter *AjudantesDataAdapter) AdicionarAjudante(firstName, lastName string, dateHired time.Time) int {
	// como o receiver deste metodo é um ponteiro ao alterarmos o nextID ele altera seu estado
	adapter.nextID++
	adapter.ajudantes[adapter.nextID] = Ajudante{
		ID:          adapter.nextID,
		Nome:        firstName,
		Sobrenome:   lastName,
		DataChegada: dateHired,
	}
	return adapter.nextID
}

// Metodo que consulta de um Ajudante por ID
// Retorna o Ajudante e um booleano indicando se ele foi encontrado
func (adapter *AjudantesDataAdapter) BuscarAjudante(id int) (Ajudante, bool) {
	e, ok := adapter.ajudantes[id]
	return e, ok
}

// Funcao que constroi um objeto de data a partir do dia, mes e ano.
func DmyToTime(dia int, mes time.Month, ano int) time.Time {
	return time.Date(ano, mes, dia, 0, 0, 0, 0, time.UTC)
}

func main() {
	ajudantesAdapter := NovoAjudantesDataAdapter()
	gerenciarAjudantes(ajudantesAdapter)
}

func gerenciarAjudantes(ed *AjudantesDataAdapter) {
	idAjudante01 := ed.AdicionarAjudante("Tomas", "Souza", DmyToTime(11, time.May, 1988))
	idAjudante02 := ed.AdicionarAjudante("Anderson", "Silva", DmyToTime(11, time.May, 1988))

	ajudante01, existe01 := ed.BuscarAjudante(idAjudante01)
	ajudante02, existe02 := ed.BuscarAjudante(idAjudante02)

	fmt.Println("Ajudante 01 existe?", existe01)
	fmt.Println("Ajudante 01: ", ajudante01)
	fmt.Println("Ajudante 02 existe?", existe02)
	fmt.Println("Ajudante 02: ", ajudante02)

	ajudante2000, existe2000 := ed.BuscarAjudante(2000)
	fmt.Println("Ajudante 2000 existe?", existe2000)
	fmt.Println("Ajudante 2000:", ajudante2000)
}
