package main

import "fmt"

func main() {
	a := 10
	// Um simples if, as chaves sao obrigatorias
	if a > 0 {
		fmt.Println("Valor de a:", a)
	}

	// Você pode inicializar uma variavel no if e no switch
	if comentario := "Você ganhou"; a > 0 {
		fmt.Println("Como o valor é maior que dez quero lhe dizer:", comentario)
	}

}
