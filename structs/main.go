package main

import "fmt"

func main()  {
	
	data := User{"Neha","neha@gmail.com",true,30}
	fmt.Println(data)
	fmt.Printf("Details are : %v\n",data)
	fmt.Printf("Details are : %+v\n",data)
	fmt.Printf("Name is %v and email is %v",data.Name,data.Email)
}

type User struct{
	Name string
	Email string
	Status bool
	Age int
}