package main

import "fmt"

func main() {
	var b byte = 30 //declarando uma variavel como byte
	var i int
	i = int(b) // convertendo de byte para int

	fmt.Printf("Valor de i = %v", i)

	var s string
	s = string(i) // converter de integer para string enterpreta o valor do integer como code point

	fmt.Println()
	fmt.Printf("String i = %v", s) // repare que o valor 30 não se mantem

}
