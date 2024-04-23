package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("start")
	wg.Add(2)
	go Printing()
	go Printing2()
	wg.Wait()
	fmt.Println("end")
	// time.Sleep(time.Millisecond * 5)
	// name := ""
	// fmt.Println("Enter the Name ")
	// fmt.Scanf("%s", &name)
	// fmt.Println("Your name is ", name)
}
func Printing() {
	defer wg.Done()

	// time.Sleep(time.Millisecond * 15)
	for i := 1; i < 100; i++ {
		fmt.Println("Press 1", i)
	}
}

func Printing2() {
	defer wg.Done()

	// time.Sleep(time.Millisecond * 15)
	for i := 1; i < 100; i++ {
		fmt.Println("Press 2", i)
	}
}
