package main

import "fmt"

func main() {
	words := []string{"a", "bola", "peixe", "abacate",
		"tomate", "carne"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "Pequena palavra!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "Palavra com tamanho exato:", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "Palavra longa!")
		}
	}
}
