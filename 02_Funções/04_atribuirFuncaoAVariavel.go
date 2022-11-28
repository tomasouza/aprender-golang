package main

import (
	"fmt"
	"os"
	"strconv"
)

// Invoque a funcao atraves do terminal 'go run 04_atribuirFuncaoAVariavel.go 20 10'
// Invoque com menos parametros ou com uma string no lugar do inteiro para validar os erros
func main() {
	if len(os.Args) != 3 {
		fmt.Println("Dois numeros inteiros sao esperados")
		os.Exit(1)
	}
	a, err := strconv.Atoi(os.Args[1])
	//Tratando o erro retornado pela funcao Atoi
	if err != nil {
		fmt.Println("primeiro argumento invalido:", err)
		os.Exit(1)
	}
	b, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("segundo argumento invalido:", err)
		os.Exit(1)
	}

	// atribuindo uma funcao a uma variavel
	soma := somarArgumentos
	// invocando a funcao atraves da variavel
	c, err := soma(a, b)
	if err != nil {
		fmt.Println("adicao falhou:", err)
	} else {
		fmt.Println(c)
	}
}

func somarArgumentos(a, b int) (int, error) {
	return a + b, nil
}
