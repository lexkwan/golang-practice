package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	md5Inst := md5.New()
	md5Inst.Write([]byte("abcdefg"))
	result := md5Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n", result)
}
