package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"io"
)

func  main()  {
	name := os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	str := fmt.Sprintf(`
	<!Doctype html>
	<html tag="en">
	<meta charset utf-8>
	<head>
	<title>Hello World!!!</title>
	</head>
	<body>
	<h1>`+name+`</h1>
	</body>
	</html>
	`)
	nf,err := os.Create("index.html")
	if(err != nil) {
		log.Fatal("error in creating file ",err)
	}
	defer nf.Close()
	io.Copy(nf,strings.NewReader(str));
}