package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string, wg *sync.WaitGroup) {
	fmt.Println("say:", s)
	time.Sleep(1 * time.Second)
	wg.Done()
}

func main() {
	t := time.Now()

	var wg sync.WaitGroup
	wg.Add(1)
	go say("Hello", &wg)
	wg.Wait()

	fmt.Println("time:", time.Since(t))
}
