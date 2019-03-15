package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func work(ctx context.Context, n int) error {
	defer wg.Done()

	c := make(chan string)
	fmt.Println("Generating goroutine")

	go func(n int) {

		fmt.Printf("Random number is %d \n", n)
		time.Sleep(time.Second * time.Duration(n))
		fmt.Printf("I have slept %d sec...\n", n)
		c <- "Done"
	}(n)

	select {
	case <-ctx.Done():
		fmt.Println("Time up. Cancel the context")
		return ctx.Err()
	case <-c:
		fmt.Println("Job done")
	}

	return nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	for i := 1; i <= 10; i++ {
		wg.Add(1)

		go work(ctx, i)
	}

	wg.Wait()

}
