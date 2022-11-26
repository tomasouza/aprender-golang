package main

import "fmt"

func main() {
	var a int = 10 // formato de declaracao completa de uma variavel
	var b = 20     // o tipo nao e necessario, pois o tipo padrao para a literal numerica 20 = int
	var c int      // o tipo especificado, mas o valor nao, inicializando o z com valor zero (0)

	fmt.Println(a, b, c)

}
