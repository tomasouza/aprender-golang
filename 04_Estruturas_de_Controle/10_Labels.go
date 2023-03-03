package main

import "fmt"

func main() {
	samples := []string{"hello", "apple_π!"}
loop_de_fora:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue loop_de_fora
			}
		}
		fmt.Println()
	}
}
