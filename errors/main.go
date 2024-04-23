package main

import (
	"fmt"
	"errors"
)



func main()  {
	result,err := divide(5,0)
	if err != nil {
		fmt.Println("Division Error :",err)
	}else{
		fmt.Println("Result :", result)
	}
}

func divide(a,b float64)(float64,error){
	if b == 0 {
		return 0, errors.New("Division by 0")
	}
	return a/b , nil
}