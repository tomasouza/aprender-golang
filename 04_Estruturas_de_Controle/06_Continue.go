package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%7 == 0 {
			fmt.Println("Divisivel por 3 e 7:", i)
			continue
		}
		if i%3 == 0 {
			fmt.Println("Divisivel por 3:", i)
			continue
		}
		if i%7 == 0 {
			fmt.Println("Divisivel por 7:", i)
			continue
		}
		fmt.Println(i)
	}
}
