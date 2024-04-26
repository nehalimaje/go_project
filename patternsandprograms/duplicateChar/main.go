package main
import (
   "fmt"
   "strings"
)
func main() {
   // Initializing the Strings
   x := "developed"
  
    // Display the Strings
   fmt.Println("First String:", x)

   
   // Using Count Function
   test1 := strings.Count(x, "d")
   test2 := strings.Count(x, "e")
  
   
   // Diplay the Count Output
   fmt.Println("Count of 'd' in the First String:", test1)
fmt.Println("Count of 'e' in the First String:", test2)
}