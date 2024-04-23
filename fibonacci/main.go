package main

import "fmt"
func main()  {
	fibonacciSeries()
}

func fibonacciSeries()  {
	a := 0
	b := 1
	
 fmt.Println(a,b)

	for i := 2; i <=5; i++ {
		sum := a + b
		fmt.Println(sum)
		a = b
		b = sum 		
	}
}