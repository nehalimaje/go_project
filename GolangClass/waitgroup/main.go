package main

import (
	"fmt"
	"sync"
	"runtime"
)

var wg sync.WaitGroup
var count int = 0

func main() {
	fmt.Println("start",runtime.NumGoroutine())
	wg.Add(2) //count of go routines can be less or equals to used goroutines
	go Printing()
	go Printing2()
	fmt.Println("middle",runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("end",runtime.NumGoroutine())
	fmt.Println("Count",count)
}
func Printing() {
	defer wg.Done()

	// time.Sleep(time.Millisecond * 15)
	for i := 1; i < 100; i++ {
		fmt.Println("Press 1", i)
		count = count + 1
	}
}

func Printing2() {
	defer wg.Done()

	// time.Sleep(time.Millisecond * 15)
	for i := 1; i < 100; i++ {
		fmt.Println("Press 2", i)
		count = count + 1
	}
}
