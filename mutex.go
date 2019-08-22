package main

import (
	"sync"
	"time"
	"fmt"
)

func main() {
	mutex := new(sync.Mutex)
	i := 0

	go func() {
		time.Sleep(20)
		mutex.Lock()
		defer mutex.Unlock()
		i += 100
	}()

	go func() {
		time.Sleep(100)
		mutex.Lock()
		defer mutex.Unlock()
		i += 10
	}()

	time.Sleep(101)
	mutex.Lock()
	i += 1
	mutex.Unlock()

	time.Sleep(500000)
	fmt.Printf("expect = 111, result = %d\n", i)
}
