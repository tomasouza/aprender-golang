package main

import (
	"fmt"
	"sync"
)

func print_variable(value string, wg *sync.WaitGroup) {
	fmt.Println(value)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	go print_variable("Primeiro", &wg)
	go print_variable("Segundo", &wg)

	wg.Add(2) // +2 Contador no wait group
	wg.Wait()
}
