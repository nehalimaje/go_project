package main

import "fmt"

func main()  {
	//CREATE CHANNEL
	number := make(chan int)
	message := make(chan string)

	//FUNCTION CALL WITH GOROUTINE
	go channelData(number,message)

	//RETRIEVE CHANNEL DATA
	fmt.Println("Channel Data : ", <-number) // ARROW symbol to get channel data
	fmt.Println("Channel Data : ", <-message)
}

func channelData(number chan int , message chan string)  {
	//SEND DATA TO CHANNEL
	number <- 15
	message <- "Learning golang"
}