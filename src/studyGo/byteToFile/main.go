package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("../bytes.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	b := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}
	n, err := f.Write(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)
}
