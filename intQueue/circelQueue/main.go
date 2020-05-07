package main

import (
	"errors"
	"fmt"
	"os"
)

type CirCleQueue struct {
	maxSize int 
	array  [5]int
	head int
	tail int
}

func (this *CirCleQueue) Push(val int) (err error){
  if this.IsFull(){
	  return errors.New("Queue is full")
  }
  this.array[this.tail] = val
  this.tail = (this.tail+1)%this.maxSize
  return
}

func (this *CirCleQueue) Pop()(val int, err error){


	if this.IsEmpty(){
		return 0, errors.New("Queue is empty")
	}
	val = this.array[this.head]
	this.head = (this.head+1)%this.maxSize
	return 
}

func (this *CirCleQueue)ListQueue(){
	size := this.Size()
	if size == 0 {
		fmt.Println("Queue is empty")
	}


	tmp := this.head
	for i:= 0; i < size; i++{
		fmt.Printf("arr[%d]=%d\t",tmp, this.array[tmp])
		tmp = (tmp+1)%this.maxSize
	}
	fmt.Println()
}

//Is full?
func (this *CirCleQueue)IsFull() bool{
	return (this.tail+1) % this.maxSize == this.head
}

//Is empty
func (this *CirCleQueue)IsEmpty() bool{
	return this.tail == this.head
}

//How many number in the queue?
func (this *CirCleQueue)Size() int{
    return (this.tail + this.maxSize- this.head) % this.maxSize
}

func main(){
	var choice string
	var val int

	q := &CirCleQueue{
		maxSize: 5,
		head:   0,
		tail:    0,
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
			err := q.Push(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("You added %d to the queue\n", val)
			}
		case "get":
			fmt.Println("Getting value from queue")
			a, err := q.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("The number return from the queue is %d\n", a)
			}

		case "show":
			q.ListQueue()
		case "exit":
			os.Exit(0)
		}
	}
}