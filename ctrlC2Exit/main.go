package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	fmt.Println("Staring app...")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// app stopped while getting the ctrl+c singal from system
	<-stop
	fmt.Println("Stopping app...")
}
