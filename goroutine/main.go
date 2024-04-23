package main

import "fmt"

func display( message string)  {
	fmt.Println(message)
}

func main()  {
	go display("Process 1")
	display("Process 2")
	display("Process 3")
}