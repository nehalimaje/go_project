package main

import "fmt"

func main() {
	fmt.Println("start")
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < 11; i++ {
			// fmt.Println("before adding value in channel", i)
			ch <- i
			// fmt.Println("after adding value in channel", i)
		}
	}()

	for v := range ch {
		fmt.Println("value ", v)
	}
}
