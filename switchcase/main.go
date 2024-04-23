package main

import 	(
	"fmt"
	"math/rand"
	"time"
)

func main(){
	fmt.Println("Switch case in go program.")
	rand.Seed(time.Now().UnixNano())
	weekdays := rand.Intn(6) + 1
	fmt.Println("Day of week is : ",weekdays)

	switch weekdays {
	case 1 :
		fmt.Println("Today is Monday")
	case 2 :
		fmt.Println("Today is Tuesday")
	case 3 :
		fmt.Println("Today is Wednesday")
	case 4 :
		fmt.Println("Today is Thursday")
	case 5 :
		fmt.Println("Today is Friday")
		fallthrough
	case 6 :
		fmt.Println("Today is Saturday")
		fallthrough
	case 7 :
		fmt.Println("Today is Sunday")
	default :
		fmt.Println("What was that!")
	}


}