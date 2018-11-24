package main

import (
	"compress/gzip"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	// cookieJar, _ := cookiejar.New(nil)
	// client := &http.Client{Transport: tr, Jar: cookieJar}

	resp, err := client.Get("http://www.hjenglish.com/")

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	var nbody string
	if resp.StatusCode == 200 {
		switch resp.Header.Get("Content-Encoding") {
		case "gzip":
			reader, _ := gzip.NewReader(resp.Body)
			for {
				buf := make([]byte, 2048)
				n, err := reader.Read(buf)
				if err != nil && err != io.EOF {
					panic(err)
				}
				if n == 0 {
					break
				}
				nbody += string(buf)
			}
		default:
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			nbody = string(bodyBytes)
		}
	}

	fmt.Println(nbody)

}
