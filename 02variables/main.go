package main

import "fmt"

//IN GO LANG WHILE DECLARING CONST, FIRST LETTER ALWAYS SHOULD BE CAPITAL
const LoginToken string = "NEHA"

func main() {
	var username string = "NEHA"
	fmt.Println(username)
	fmt.Printf("Username is of type %T \n", username)

	var isBoolean bool = true
	fmt.Println(isBoolean)
	fmt.Printf("isBoolean is of type %T \n", isBoolean)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("small val is of type %T \n", smallVal)

	var smallfloat float64 = 255.45697463
	fmt.Println(smallfloat)
	fmt.Printf("smallFloat is of type %T \n", smallfloat)

	//default values
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("anotherVariable is of type %T \n", anotherVariable)

	//implicit type
	var website = "learncodeonline.in" //if you not define data type it will automaticaly takes it
	fmt.Println(website)
	fmt.Printf("website is of type %T \n", website)

	// no var style
	noOfUsers := 300000 // we can only use this function ( := ) inside function, outside fun its not allowed
	fmt.Println(noOfUsers)
	fmt.Printf("noOfUser is of type %T \n", noOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Login token is of type %T \n", LoginToken)

}
