package main

import (
	"sync"
	"time"
	"fmt"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)

	go func() {
		time.Sleep(10000000)
		fmt.Println("ahooooooooooooooooooooo")
		wg.Done()
	}()

	wg.Wait()
	fmt.Printf("おわたーーーーー")
}
