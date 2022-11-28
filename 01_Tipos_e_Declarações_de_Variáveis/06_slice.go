package main

import "fmt"

func main() {
	//slices sao listas sem tamanho definido eles podem crescer
	var intSlice []int                                      //declarando um slice de tipo int vazio
	stringSlice := []string{"the", "quick", "brown", "fox"} //declarando um slice de tipo string com valores iniciais
	zeroIntSlice := make([]int, 10)                         // declarando um slice com 10 variaveis do tipo int de valor padrao

	fmt.Println("valor de intSlice:", intSlice)
	fmt.Println("tamanho de intSlice:", len(intSlice))

	fmt.Println("valor de stringSlice:", stringSlice)
	fmt.Println("tamanho de stringSlice:", len(stringSlice))

	fmt.Println("valor de zeroIntSlice:", zeroIntSlice)
	fmt.Println("tamanho de zeroIntSlice:", len(zeroIntSlice))

	//podemps recuperar um item do slice atraves de seu indice
	fmt.Println("primeiro e ultimo valor de stringSlice:", stringSlice[0], stringSlice[len(stringSlice)-1])

	// atribuindo valor ao index 3 da stringSlice
	stringSlice[3] = "bear"
	fmt.Println("novo ultimo valor de stringSlice:", stringSlice[len(stringSlice)-1])

	// adicionando na ultima posicao mais um elemento a intSlice
	intSlice = append(intSlice, 20)
	fmt.Println("novo valor de intSlice:", intSlice)

	// adicionando na ultima posicao mais um elemento a zeroIntSlice
	zeroIntSlice = append(zeroIntSlice, 100)
	fmt.Println("novo valor de zeroIntSlice:", zeroIntSlice)

	// demonstrando que os slices aumentaram depois das adicoes de novos elementos
	fmt.Println("novos tamanhos de intSlice s stringSlice", len(intSlice), len(zeroIntSlice))

	// podemos extrair um slice de um outro slice
	midSlice := stringSlice[1:3]
	startSlice := stringSlice[:2]
	endSlice := stringSlice[2:]
	fmt.Println("valor de midSlice:", midSlice)
	fmt.Println("valor de startSlice:", startSlice)
	fmt.Println("valor de endSlice:", endSlice)
}
