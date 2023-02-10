package main

import (
	"fmt"
	"sync"
	"time"
)

func def_value(v chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	v <- 5
}

func main() {
	var wg sync.WaitGroup

	ch := make(chan int)

	wg.Add(1)
	go def_value(ch, &wg)
	fmt.Println("Waiting for a value")
	fmt.Printf("Value is %d\n", <-ch)

	wg.Wait()
}
