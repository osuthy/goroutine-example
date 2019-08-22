package main

import (
	"sync"
	"time"
	"fmt"
)

func main() {
	mutexY := new(sync.Mutex)
	mutexZ := new(sync.Mutex)

	i := 0

	go func() {
		mutexY.Lock()
		time.Sleep(20)
		mutexZ.Lock()
		i += 100
		defer mutexY.Unlock()
		defer mutexZ.Unlock()

	}()

		mutexZ.Lock()
		time.Sleep(100)
		mutexY.Lock()
		i += 10
		defer mutexZ.Unlock()
		defer mutexY.Unlock()


	time.Sleep(500000)
	fmt.Printf("expect = 111, result = %d\n", i)
}
