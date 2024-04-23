package main

import "fmt"

func main()  {
	
	data := User{"Neha","neha@gmail.com",true,30}
	fmt.Println(data)
	fmt.Printf("Details are : %v\n",data)
	fmt.Printf("Details are : %+v\n",data)
	fmt.Printf("Name is %v and email is %v\n",data.Name,data.Email)
	data.getStatus();
	data.getEmail();
	fmt.Printf("Email is %v\n",data.Email)
	data.getUpdatedEmail()
	fmt.Printf("Email is %v\n",data.Email)
}

type User struct{
	Name string
	Email string
	Status bool
	Age int
}

func (u User) getStatus(){
	fmt.Println("Status of this user is :",u.Status)
}

func (u User) getEmail(){
	u.Email = "test@gmail.com"
	fmt.Println("New Email is :", u.Email)
}
func (u *User) getUpdatedEmail(){
	u.Email = "test@gmail.com"
	fmt.Println("Updated Email with Pointer..")
}