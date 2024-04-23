package main

import "fmt"

type Book struct{
	name string
	author string
	pages int
}

type Shelf struct{
	position int
	book Book
}

func structdemo(){
	book := Book{name : "Harry Potter",author : "J.K.Roling",pages : 800}
	fmt.Println(book)
	fmt.Println("Name : ",book.name)
	fmt.Println("Author : ",book.author)
	fmt.Println("Pages : ",book.pages)
	fmt.Println()

	fmt.Println("Book Shelf details.")
	s := Shelf{position : 8, book : book}
	fmt.Println(s)
	fmt.Println("Book Position : ",s.position)
	fmt.Println("Book details : ",s.book)
	fmt.Println()
}