package main

import (
	"fmt"
)

func main() {
	//Agora buscando um número de 1 a 5 mas com blank switch
	for i := 0; i < 10; i++ {
		switch {
		case i == 0:
			fmt.Println("Este número é muito baixo:", i)
		case i > 5:
			fmt.Println("Este número é muito alto", i)
		default:
			fmt.Println("Este é um bom número:", i)
		}
	}
}
