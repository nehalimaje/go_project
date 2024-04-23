package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// welcome := "Welcome to user input"
	// fmt.Println(welcome)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter rating for our pizza : ")

	// //comma ok || err
	// input, _ := reader.ReadString('\n') // store user input in input variable
	// fmt.Println("Thanks of your rating ", input)
	// fmt.Printf("Type of input is %T \n", input)

	enterfirstname := "Please enter your first name"
	fmt.Println(enterfirstname)

	fname := bufio.NewReader(os.Stdin)
	fnameInput, _ := fname.ReadString('\n')

	enterlastname := "Please enter your last name"
	fmt.Println(enterlastname)

	lname := bufio.NewReader(os.Stdin)
	lnameInput, _ := lname.ReadString('\n')

	fmt.Println("Your first name is : ", fnameInput)
	fmt.Println("Your last name is : ", lnameInput)

}
