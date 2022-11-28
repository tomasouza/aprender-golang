package main

import (
	"fmt"
	"os"
)

// Execute a funcao no terminal e varie a quantidade de parametros 'go run 03_validarEntrada.go'
func main() {
	/*Repare que estamos validando se a funcao foi invocada com tres argumentos
	O primeiro seria o nome da funcao e os outros dois as entradas que esperamos*/
	if len(os.Args) != 3 {
		fmt.Println("Dois parametros sao esperados")
		os.Exit(1)
	}

	fmt.Println("Voce acertou a quantidade de parametros esperados")
	fmt.Println("Parametro um:", os.Args[1])
	fmt.Println("Parametro dois:", os.Args[2])
}
