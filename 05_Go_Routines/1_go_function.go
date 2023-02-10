package main

import (
	"fmt"
)

func print_var(value string) {
	fmt.Println(value)
}

func main() {
	go print_var("Primeiro")
	go print_var("Segundo")

	fmt.Println("Não imprimiu nada") // Comentar e descomentar a debaixo
	//time.Sleep(2 * time.Second)
}
