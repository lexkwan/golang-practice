package main

func binrayFind(arr []int, start int, end int, val int) {
	if start > end {
		println("Can't be found")
		return
	}
	middle := (end + start) / 2
	if arr[middle] < val {
		binrayFind(arr, middle+1, end, val)
	} else if arr[middle] > val {
		binrayFind(arr, start, middle-1, val)
	} else {
		println("Found the number!!")
	}
}

func main() {
	arr := []int{1, 3, 7, 66, 99}
	binrayFind(arr, 0, len(arr)-1, 3)

}
