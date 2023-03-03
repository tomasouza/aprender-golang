package main

import "fmt"

func main() {
	//Estamos procurando por uma palavra de tamanho 5 neste caso
	words := []string{"a", "bola", "peixe", "abacate",
		"eletrecista", "carne"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "é uma palavra pequena!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "é uma palavra com tamanho exato:", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "é uma palavra longa!")
		}
	}
}
