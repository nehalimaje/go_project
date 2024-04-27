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

findDuplicateChar()
}

func findDuplicateChar(){
   str := "developed"
   duplicateCount := 0

    for i := 0; i < len(str); i++ {
        count := 1
        // Check if the character at index i has already been counted
        if str[i] != 0 {
            for j := i + 1; j < len(str); j++ {
                if str[i] == str[j] {
                    count++
                    // Set the character at index j to 0 to mark it as counted
              
                    str[j] = 0
                }
            }
            // If count is greater than 1, it means the character is a duplicate
            if count > 1 {
                duplicateCount++
                fmt.Printf("Character '%c' has %d duplicates\n", str[i], count)
            }
        }
    }

}