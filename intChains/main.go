package main

import (
	"fmt"
)

type IntChain struct {
	val  int
	next *IntChain
}

func (c *IntChain) Next() int {
	return c.next.val
}

func main() {
	var a, b, c, d, e, f IntChain

	a = IntChain{val: 1, next: &b}
	b = IntChain{val: 10, next: &c}
	c = IntChain{val: 9, next: &d}
	d = IntChain{val: 12, next: &e}
	e = IntChain{val: 3, next: &f}
	f = IntChain{val: 4, next: &a}

	pointer := a
	ss := []int{1, 2, 3, 4, 5, 6}
	for _, i := range ss {
		fmt.Printf("%d * %d \n", i, pointer.val)
		pointer = *pointer.next
	}
}
