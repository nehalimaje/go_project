package main

import (
	"fmt"
	"log"
	"mongoapi/router"
	"net/http"
)

func main() {
	fmt.Println("Mongo DB API")
	r := router.Router()
	fmt.Println("Server is getting started....")
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Server is listening at port 8080 ...")
}
