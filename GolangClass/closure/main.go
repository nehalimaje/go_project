package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	max := 51
	var mp = map[string]string{}
	mp = map[string]string{"AC": "AC", "BC": "BC", "CC": "Value"}
	fmt.Println("", mp)
	listOfKeys := []string{}
	for Key, Value := range mp {
		fmt.Println("Key , value is ", Key, Value)
		listOfKeys = append(listOfKeys, Key)

	}
	fmt.Println("List Of Keys ", listOfKeys)
	for i := 0; i < 10; i++ {
		if i == max {

			fmt.Println("MATCH ", max, i)
			break

		}

		first, second := Test()
		fmt.Println("", i, first, second)
		vartestclose := TestClose()
		fmt.Println("vartestclose", vartestclose())
	}

	testFuncVar := ClousreFunction()

	fmt.Println(" Clousre Function  ", testFuncVar())
}

func Test() (first, second int) {

	if first == 0 {
		fmt.Println("test")

	}
	return 2, 3

}

func ClousreFunction() func() int {
	Test()
	fmt.Println("test outer")
	return func() int {
		fmt.Println("Inner")
		return 12
	}
}

func TestClose() func() float32 {

	return func() float32 {
		return 34.4
	}
}
