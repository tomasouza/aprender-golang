package main

import (
	"fmt"
	"os"
	"strconv"
)

// Execute a funcao no terminal passando um inteiro de 1 a 3 para escolher a funcao a rodar
// 'go run 06_decidirQualFuncaoInvocar.go 1'
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Um parametro numerico esperado")
		os.Exit(1)
	}
	val, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Parametro invalido:", err)
		os.Exit(1)
	}
	//decalara a variavel que contera a funcao escolhida
	var funcaoEscolhida func()
	switch val {
	case 1:
		funcaoEscolhida = Um
	case 2:
		funcaoEscolhida = Dois
	case 3:
		funcaoEscolhida = Tres
	default:
		fmt.Println("Funcao inexistente, favor escolher de 1 a 3")
		os.Exit(1)
	}
	funcaoEscolhida()
}

func Um() {
	fmt.Println("Funcao Um invocada")
}
func Dois() {
	fmt.Println("Funcao Dois invocada")
}
func Tres() {
	fmt.Println("Funcao Tres invocada")
}
