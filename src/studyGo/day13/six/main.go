package main

import "fmt"

func writeData(c chan int) {
	for i := 1; i <= 50; i++ {
		fmt.Println("write", i)
		c <- i
	}
	close(c)
}

func readData(c chan int, e chan bool) {
	for {
		v, ok := <-c
		fmt.Println("read", v)
		if !ok {
			break
		}
	}
	e <- true
	close(e)
}
func main() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, exitChan)
	go readData(intChan, exitChan)
	go readData(intChan, exitChan)
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
