package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("", reflect.TypeOf(GetErr))
	fmt.Println("")

	defer func() {
		if msg := GetErr(); msg != nil {
			fmt.Println("HANDLE PANIC ")
		}
	}()

}

func GetErr() error {

	fmt.Println("")
	for i := 0; i < 11; i++ {
		if i == 5 {
			panic("PANIC IS CREATED")

		}
	}
	return errors.New("TEST ERROR IS CREATED")
}
