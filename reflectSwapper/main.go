package main

import (
	"fmt"
	"reflect"
)

func main()  {
	theList := []int{1,2,3,4,5}
	swap := reflect.Swapper(theList)
	fmt.Println("Original slice :",theList)
	swap(2,3)
	fmt.Println("After Swapped :",theList)
	// Reversing a slice using Swapper() function
	for i := 0; i < len(theList)/2; i++ {
		swap(i, len(theList)-1-i)
	}
	fmt.Printf("After Reverse :%v\n", theList)
}