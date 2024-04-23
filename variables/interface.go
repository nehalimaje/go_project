package main

import "fmt"

type userProfile interface {
	getName() string
	getAddress() string
	getEmail() string
}

func getUserInfo(u userProfile) {
	fmt.Println("Name ",u.getName())
	fmt.Println("Address ", u.getAddress())
	fmt.Println("Email ",u.getEmail())
}

// func interfaceDemo(){
// 	s := userProfile{
// 		Name : "Sherlock Holmes",
// 		Address : "2218, Baker Street",
// 		Email : "sherlock@homes.com",
// 	}

//}
//getUserInfo();