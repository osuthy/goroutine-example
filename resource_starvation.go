package main

import (
	"sync"
	"time"
	"fmt"
)

func main() {
	mutex := new(sync.Mutex)

	go func() {
		mutex.Lock()
		fmt.Println("ahoooooooooo")
	}()

	time.Sleep(100000)
	mutex.Lock()
	fmt.Println("ahoooooooooo")
}
