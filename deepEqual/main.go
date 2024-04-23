package main

import (
	"fmt"
	"reflect"
)

func main()  {
	s1 := []string{"a","b","c"}
	s2 := []string{"c","d"}
	stringResult := reflect.DeepEqual(s1,s2)
	fmt.Println(stringResult)

	n1 := [3]int{1,2,3}
	n2 := [3]int{1,2,3}
	numResult := reflect.DeepEqual(n1,n2)
	fmt.Println(numResult)

	a := [3]int{1,2,3}
	b := [3]int{4,5,6}
	result := reflect.DeepEqual(a,b)
	fmt.Println(result)
}

