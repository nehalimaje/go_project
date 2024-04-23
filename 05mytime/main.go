package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golan")

	presentTime := time.Now()
	fmt.Println("Current time is ", presentTime)
	fmt.Println("current date is ", presentTime.Format("2006-01-02"))
	fmt.Println("current time is ", presentTime.Format("15-04-05 Wednesday"))

	createDate := time.Date(2022, time.February, 07, 23, 23, 0, 0, time.UTC)
	fmt.Println(createDate)
	fmt.Println("created date is ", createDate.Format("2006-01-02 Monday"))
}
