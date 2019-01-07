package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	dirPath := "abc/def/ghi"

	if _, err := os.Stat(dirPath); !os.IsNotExist(err) {
		fmt.Println(dirPath + " is exist.")
		fmt.Println("Removing this folder")
		os.RemoveAll(dirPath) // would remove the last folder ghi
	}

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.MkdirAll(dirPath, 0766)
	}

	if _, err := os.Stat(dirPath); !os.IsNotExist(err) {
		fmt.Println(dirPath + " Created successfully.")
	}

	fmt.Println("Creating tmp.txt file")
	ioutil.WriteFile(dirPath+"/tmp.txt", []byte("This a test file."), 660)

}
