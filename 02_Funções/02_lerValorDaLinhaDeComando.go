package main

import (
	"fmt"
	"os"
	"strconv"
)

// Para executar esta funcao digite no seu terminal 'go run 02_lerValorDaLinhaDeComando.go 20 10'

func main() {
	/* Este _ referencia uma variavel da qual nao gostariamos de utilizar
	Por padrao o compilador exige que voce referencie uma variavel inutilizada com _ */
	argumentoUm, _ := strconv.Atoi(os.Args[1]) // A funcao Atoi converte de string para int
	// a fincao Atoi retorna (int, error) no momento estamos decidindo ignorar o erro
	argumentoDois, _ := strconv.Atoi(os.Args[2])

	fmt.Println("Primeiro argumento:", argumentoUm)
	fmt.Println("Segundo argumento", argumentoDois)

}
