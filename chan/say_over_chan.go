package main

import (
	"fmt"
	"math/rand"
	"time"
)

func say(msg string, c chan string) {
	for i := 0; ; i++ {
		v := fmt.Sprintf("%s %d", msg, i)
		c <- v
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	c := make(chan string)

	go say("Gopher!", c)

	for {
		msg := <-c
		fmt.Println(msg)
	}
}
