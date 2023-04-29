package main

import "fmt"

func main() {
	c := make(chan string, 2)
	c <- "gopher"
	c <- "panda"

	m1 := <-c
	fmt.Println("m1:", m1)
	m2 := <-c
	fmt.Println("m2:", m2)

	c <- "xxx"

	m3 := <-c
	fmt.Println("m3:", m3)
}
