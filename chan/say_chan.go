package main

import "fmt"

func main() {
	c := make(chan string)

	c <- "Hello"

	msg := <-c

	fmt.Println("msg:", msg)
}
