package main

import (
	"fmt"
)

var ch = make(chan int, 5)

func main() {
	fmt.Println("start")
	go func() {
		defer close(ch)
		for i := 0; i < 11; i++ {
			ch <- i
			// fmt.Println("afterr")
		}

	}()
	for v := range ch {
		fmt.Println(" ", v)
	}
	// time.Sleep(time.Millisecond * 6)
}
