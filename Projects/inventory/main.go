package main

import (
	"fmt"
	"inventory/router"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Welcome to Inventory Project")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Server is listening at port 8080 ...")
}
