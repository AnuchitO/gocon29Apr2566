package main

import (
	"fmt"
	"math/rand"
	"time"
)

func say(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}

func main() {
	gopher := say("gopher")
	panda := say("panda")

	for {
		select {
		case msg := <-gopher:
			fmt.Println(msg)
		case msg := <-panda:
			fmt.Println(msg)
		}
	}
}
