package main

import "fmt"

func main(){
myNumber := 23

var ptr = &myNumber

fmt.Println("Address of pointer is ",ptr)

fmt.Println("Value of pointer is ",*ptr)

*ptr = *ptr + 2

fmt.Println("Value of pointer is ",myNumber)
}