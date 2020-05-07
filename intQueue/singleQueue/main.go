package main

import (
	"errors"
	"fmt"
	"os"
)

type Queue struct {
	maxSize int
	array   [5]int
	front   int
	rear    int
}

func (this *Queue) addQueue(val int) error {
	// is it Queue full?
	if this.rear == this.maxSize-1 {
		return errors.New("Sorry, the queue is full")
	}
	this.rear++
	this.array[this.rear] = val
	return nil
}

func (this *Queue) getQueue() (val int, err error) {
	if this.front == this.rear {
		return -1, errors.New("Empty queue")
	}
	this.front++
	val = this.array[this.front]
	return val, nil
}

func (this *Queue) showQueue() {
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d]=%d ", i, this.array[i])
	}
	fmt.Println()
}

func main() {
	var choice string
	var val int

	q := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
	}
	for {
		fmt.Println("1.add")
		fmt.Println("2.get")
		fmt.Println("3.show")
		fmt.Println("4.exit")

		fmt.Scanln(&choice)

		switch choice {
		case "add":
			fmt.Println("Please input a value you want to add:")
			fmt.Scanln(&val)
			err := q.addQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("You added %d to the queue\n", val)
			}
		case "get":
			fmt.Println("Getting value from queue")
			a, err := q.getQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("The number return from the queue is %d\n", a)
			}

		case "show":
			q.showQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
