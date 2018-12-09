package main

import "fmt"

func reverseINT(n int) int {
	newInt := 0
	for n > 0 {
		remainder := n % 10
		newInt = newInt*10 + remainder
		n /= 10
	}
	return newInt
}

func main() {
	fmt.Printf("Reversing int 123456: %d\n", reverseINT(123456))
	fmt.Printf("Reversing int 123321: %d\n", reverseINT(123321))
}
