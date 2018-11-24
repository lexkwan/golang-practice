package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`(.*?)\(`)
	rs := re.FindStringSubmatch("sys/disk(456,789)")
	fmt.Println(re[1])
}
