package main

import (
	"time"
	"fmt"
)

func main() {
	i := 0

	go func() {
		time.Sleep(20)
		i += 100
	}()

	go func() {
		time.Sleep(100)
		i += 10
	}()

	time.Sleep(101)
	i += 1

	time.Sleep(500000)
	fmt.Printf("expect = 111, result = %d\n", i)
}
