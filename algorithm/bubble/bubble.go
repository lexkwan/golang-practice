package main

import (
	"fmt"
)

func bubble(arr *[5]int) {
	fmt.Printf("Before sorted: array=%v\n", *(arr))
	var temp int

	for x := 0; x < len(arr)-1-x; x++ {
		for y := 0; y < len(arr)-1; y++ {
			if arr[y] > arr[y+1] {
				temp = arr[y]
				arr[y] = arr[y+1]
				arr[y+1] = temp
			}
		}
	}

	fmt.Printf("After sorted: array=%v\n", *(arr))
}

func main() {
	arr := [5]int{10, 33, 4, 99, 66}
	bubble(&arr)
}
