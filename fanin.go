package main

import (
	"fmt"
	"time"
)

func say(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(1e3) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() { for { c <- <-input1 } }()
	go func() { for { c <- <-input2 } }()
	return c
}

func main() {
	gop := say("gopher")
	pen := say("penguin")
	c := fanIn(gop, pen)
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("gopher what's up.")
}
