package main

import "fmt"

func main(){
	fmt.Println("Welcome to the loop in go programming.")
	fruits := []string{"Apple","Mango","Banana","Orange"}
	fmt.Println(fruits)

	for f := 0; f < len(fruits); f++ {
		fmt.Println(fruits[f])
	}

	for i := range fruits{
		fmt.Println("Using range",fruits[i])
	}

	for index,fruit := range fruits {
		fmt.Printf("Index is %v and fruit is %v\n",index , fruit)
	}

	num := 1

	for num < 10 {
		if num == 3 {
			num++
			continue
		}
		// if num == 8 {
		// 	break
		// }
		if num == 7 {
			goto lco
		}
		fmt.Println("value is ",num)
		num++
	}

	lco :
	fmt.Println("Welcome to goto statment of go programming")
}