package main

import "fmt"

func main() {
	var a int = 10
	var b = 20
	var c int

	fmt.Printf("Valores declarados: %v %v %v", a, b, c)

	a = b      // atribuindo o valor de 'b' a 'a'
	b = 30     // atribuindo um novo valor a b
	c = 3 * 50 // multiplicando 3 vezes 50 e atribuindo a 'c'

	fmt.Println()
	fmt.Printf("Valores atribuidos posteriormente: %v %v %v", a, b, c)

}
