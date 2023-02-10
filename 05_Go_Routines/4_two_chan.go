package main

import (
	"fmt"
	"sync"
)

const MULTIPLE = 2

func get_value(v chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	v <- 5
}

func calculate(v chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	a := <-v
	fmt.Printf("Value is %d\n", a*MULTIPLE)
}

func main() {
	var wg sync.WaitGroup

	ch := make(chan int)

	go get_value(ch, &wg)
	go calculate(ch, &wg)

	wg.Add(2)
	wg.Wait()
}
