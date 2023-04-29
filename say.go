package main

import (
	"fmt"
	"time"
)

func say(s string) {
	fmt.Println("say:", s)
	time.Sleep(1 * time.Second)
}

func main() {
	t := time.Now()

	go say("Hello")

	fmt.Println("time:", time.Since(t))
}
