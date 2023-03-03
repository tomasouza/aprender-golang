package main

import "fmt"

func main() {
	//Observar que os valores entregues pelo for range sao copias
	valoresPares := []int{2, 4, 6, 8, 10, 12}
	for _, v := range valoresPares {
		v *= 2
	}
	//Ele nao altera o valor do slice
	fmt.Println(valoresPares)
	x := 2
	x *= 2
	fmt.Println(x)
}
