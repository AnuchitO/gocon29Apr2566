package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("say:", s)
	time.Sleep(1 * time.Second)
}

func main() {
	t := time.Now()

	var wg sync.WaitGroup
	wg.Add(1)
	go say("Hello", &wg)
	wg.Wait()

	fmt.Println("time:", time.Since(t))
}
