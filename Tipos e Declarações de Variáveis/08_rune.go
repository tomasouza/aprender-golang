package main

import "fmt"

func main() {
	//Go tem um tipo chamado rune para representar um Unicode code point.
	//Um rune é um int de 32-bit e seu literal um unico character com aspas simples.
	baseString := "The quick brown fox jumped over the lazy dog"
	stringRunes := []rune(baseString) // convertendo uma string em um slice de runes
	fmt.Println(len(stringRunes), stringRunes)

	roundTripRunes := string(stringRunes) // convertendo novamente de runes a string
	fmt.Println(roundTripRunes)
	// declarando um literal de rune
	r := 'X'
	fmt.Println(r, string(r))
}
