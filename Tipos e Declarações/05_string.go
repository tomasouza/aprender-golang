package main

import "fmt"

func main() {
	s1 := "Esta é uma literal de string interpretada.\n\tEla precisa ser escrita em uma unica linha.\n\tEscape sequences sao interpretadas."
	fmt.Println(s1)

	s2 := `Esta é uma literal de string raw.
	Ela pode ser escrita entre multiplas linhas.
	Escape sequences como \n nao sao interpretadas.`
	fmt.Println(s2)
}
