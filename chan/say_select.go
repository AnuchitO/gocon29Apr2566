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
	c1 := make(chan string, 5)
	c2 := make(chan string, 5)

	go say("Gopher!", c1)
	go say("Panda!", c2)

	for {
		select {
		case m := <-c1:
			fmt.Println(m)
		case msg := <-c2:
			fmt.Println(msg)
		default:
			fmt.Println("Nothing ready")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
