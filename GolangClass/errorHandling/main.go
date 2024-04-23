package main

import (
	"errors"
	"fmt"
)

type CustomErr struct {
	Data    Emp
	Message error
}

type Emp struct {
	Name string
	Id   int
}

func recoverDemo() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	panic("Something went wrong!")
}

func main() {
	fmt.Println("Start")
	recoverDemo()
	// AddData()
	fmt.Println("AddData", AddData())
	fmt.Println("End")
}

func AddData() CustomErr {
	emp := Emp{
		Name: "Error", Id: 132,
	}
	mess := CustomErr{
		emp, errors.New("test err"),
	}
	return mess
}
