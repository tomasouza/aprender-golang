package main

import "fmt"

func main() {
	//Strings possuem capacidades similares a slices
	baseString := "The quick brown fox jumped over the lazy dog"
	fmt.Println("valores em byte do char na string:", baseString[0], baseString[len(baseString)-10])
	fmt.Println("convertendo para string os bytes:", string(baseString[0]), string(baseString[len(baseString)-10]))
	fmt.Println("substring de baseString:", baseString[16:26])

	sunString := "It's a 🌞 morning!"                     // parece que o tamnaho deveria ser de 17...
	fmt.Println("tamanho de sunString:", len(sunString)) // como o tamanho do emoji é maior do que um byte ele exibe 20

	//convertendo strings e slices
	stringBytes := []byte(baseString)
	fmt.Println("tamanho de stringBytes:", len(stringBytes), stringBytes)
	roundTrip := string(stringBytes)
	fmt.Println(roundTrip)
	b := sunString[0]
	fmt.Println(b, string(b))
}
