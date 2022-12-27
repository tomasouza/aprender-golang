package main

import "fmt"

// Utilizar o pacote FMT para formatar a saída - ref: https://pkg.go.dev/fmt
func main() {
	var numero int = 14
	var texto string = "TEXTO"
	var booleano bool = true
	var pontoFlutuante float64 = 42.42424242

	fmt.Printf("Valores numero (base 10): %d \n", numero)
	fmt.Printf("Valores numero (base 2 - binário): %b \n", numero)
	fmt.Printf("Valores numero (base 16 - hexadecimal): %x \n", numero)
	fmt.Printf("Valores texto: %s \n", texto)
	fmt.Printf("Valores booleano: %t \n", booleano)
	fmt.Printf("Valores ponto flutuante (2 casas): %.2f \n", pontoFlutuante)
}
