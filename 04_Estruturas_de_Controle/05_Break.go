package main

import "fmt"

func main() {
	x := 0
	for {
		x++
		fmt.Println("Iteracao numero", x)
		if x > 9 {
			break
		}
	}

}
