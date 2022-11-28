package main

import "fmt"

func main() {
	a := 10 // o operador := declara e atribui valor a uma variavel
	b := 20
	fmt.Printf("Valores de a e b: %v %v", a, b)

	a, b, c, d := 20, 40, 60, 80 // você pode declar mais do que uma variavel com :=

	fmt.Println()
	fmt.Printf("Valores de a, b, c, d: %v %v %v %v", a, b, c, d)
}
