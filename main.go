package main

import "fmt"

//! Channels in Go

func sendValues(c chan int) {
	c <- 9 //* send opertaor
}

func main() {
	values := make(chan int)

	defer close(values)

	go sendValues(values)

	value := <-values
	fmt.Println(value)
}
