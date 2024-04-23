package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup
var count int = 0

func main() {
	fmt.Println("start", runtime.NumGoroutine())
	wg.Add(2)
	go Printing()
	go Printing2()
	fmt.Println("Middle ", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("End", runtime.NumGoroutine())
	fmt.Println("end", count)
	// time.Sleep(time.Millisecond * 5)
	// name := ""
	// fmt.Println("Enter the Name ")
	// fmt.Scanf("%s", &name)
	// fmt.Println("Your name is ", name)
}
func Printing() {
	defer wg.Done()

	for i := 1; i <= 100; i++ {
		temp := count //100     //actual 101
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
		count = temp + 1
		fmt.Println("Press 1", i)
	}
}

func Printing2() {
	defer wg.Done()

	for i := 1; i <= 100; i++ {
		temp := count
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(5)))
		count = temp + 1
		fmt.Println("Press 2", i)
	}
}
