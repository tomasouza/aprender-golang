package main

import "fmt"

func main() {
	// quando estamos falando de iterar sobre strings, ele itera sobre cada um dos runes
	samples := []string{"hello", "apple_π!"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}
	//Algo surpreendente que acontece e que alguns runes necessitam de mais de um byte
}
