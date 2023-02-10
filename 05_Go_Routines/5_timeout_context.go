package main

import (
	"context"
	"fmt"
	"time"
)

func print_timeout(ctx context.Context, t int) {
	defer func() {
		fmt.Printf("Bye bye -> %d timeout\n", t)
	}()
	fmt.Printf("Sleeping for %d seconds\n", t)
	time.Sleep(time.Duration(t) * time.Second)
	fmt.Printf("Sleep occured ->%d<- with success\n", t)
}

func main() {
	// Create a context with a timeout of 5 seconds
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*7)
	defer cancel()

	go print_timeout(ctxTimeout, 5)
	go print_timeout(ctxTimeout, 10)

	select {
	case <-ctxTimeout.Done():
		fmt.Println(ctxTimeout.Err())
		fmt.Println("cancelling...")
		cancel()
	}
	time.Sleep(5 * time.Second)
}
