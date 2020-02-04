package main

import (
	"fmt"
	"strings"
)

// makeSuffix closer funtion example
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// makeSuffix2 non-closer fuction solution
func makeSuffix2(suffix string, name string) string {
	if !strings.HasSuffix(name, suffix) {
		return name + suffix
	}
	return name
}

func main() {
	f := makeSuffix(".jpg")
	fmt.Println(f("lex"))
	fmt.Println(f("kwan.jpg"))
	fmt.Println(makeSuffix2(".jpg", "lex"))
	fmt.Println(makeSuffix2(".jpg", "kwan.jpg"))
}
