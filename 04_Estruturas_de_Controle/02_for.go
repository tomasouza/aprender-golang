package main

import "fmt"

func main() {
	// você pode declarar um indice no loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Somatório dos índices:", sum)

	mapa := make(map[string]string)
	mapa["chave"] = "valor"
	mapa["corvete"] = "rua 33"
	mapa["opala"] = "rua 12"

	// se você quiser iterar sobre um array, slice, string ou mapa
	// ou ler de um channel temos o range
	for chave, valor := range mapa {
		fmt.Println("chave:", chave)
		fmt.Println("valor:", valor)
	}

	//se você precisar somente do primeiro valor
	for chave := range mapa {
		fmt.Println("Chave:", chave)
	}

	//se você precisar somente do segundo valor
	for _, valor := range mapa {
		fmt.Println("Valor:", valor)
	}

}
