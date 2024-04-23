package main

import "fmt"
type profile struct {
	name string
	address string
	email string
	contact string
}

func (p profile) getname() string{
	return p.name
}

func (p profile) getaddress() string{
	return p.address
}

func (p profile) getemail() string{
	return p.email
}

func (p profile) getcontact() map[string]string {
	m := map[string]string{"email ":p.email,"contact ":p.contact}
	return m
}

func methodDemo() {
s := profile{
	name : "Sherlock",
	address : "2218, Baker Street",
	email : "sherlock@homes.com",
	contact: "1888956",
}

fmt.Println("Name ",s.getname())
fmt.Println("Address ",s.getaddress())
fmt.Println("Contact ",s.getcontact())
fmt.Println("Email ",s.getcontact()["email"])

}
