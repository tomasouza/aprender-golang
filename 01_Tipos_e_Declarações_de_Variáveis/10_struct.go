package main

import "time"

// Podemos declarar nosso proprio tipo com os structs
type Estudante struct {
	ID              int
	Nome            string
	Sobrenome       string
	DataDeMatricula time.Time
}
