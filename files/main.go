package main

import (
	"fmt"
	"os"
	"io"
	//"string"
	"io/ioutil"
)

func main(){
	fmt.Println("Welcome to files in golang")
	content := "This is sample content"

	file,err := os.Create("./golangFile.txt")

	if err != nil {
		panic(err)
	}

	length,err := io.WriteString(file,content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is ",length)
	defer file.Close()

	readFile("./golangFile.txt")	
}

func readFile(filename string){
	databyte,err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Content of file is :",string(databyte))
}