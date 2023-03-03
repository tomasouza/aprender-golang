package main

import "fmt"

func main() {
	//Ao declarar uma variável dentro do bloco main ela pertencerá a este bloco
	x := 6
	if x > 3 {
		fmt.Println(x)
		//Ao declarar esta segunda variável de mesmo nome, ela pertencerá a este bloco if
		x := 12
		fmt.Println(x)
	}
	fmt.Println(x)
}
