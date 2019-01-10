package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	t := time.Now().Add(100 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), t)

	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("time up.")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

}
