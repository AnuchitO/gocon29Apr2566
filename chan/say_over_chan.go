package main

import (
	"fmt"
	"math/rand"
	"time"
)

func say(msg string) {
	for i := 0; ; i++ {
		v := fmt.Sprintf("%s %d", msg, i)
		fmt.Println("say:", v)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	c := make(chan string)
	_ = c

	go say("Gopher!")

	fmt.Scanln()
}
