package main

import(
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main()  {
	router := gin.Default()
	router.Run(":8080")
}