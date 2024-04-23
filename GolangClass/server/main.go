// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	fmt.Println("start")
// 	http.HandleFunc("/gettwonumber", sumHandler)
// 	fmt.Println("Server started on 9090 port")
// 	http.ListenAndServe(":9090", nil)
// }

// func sumHandler(rw http.ResponseWriter, rq *http.Request) {
// 	rw.Write([]byte("Server is Readyyy"))
// }

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Numbers struct {
	Num1 int `json:"num1" bson:"number"`
	Num2 int `json:"num2"`
}

type SumResponse struct {
	Sum int `json:"sum"`
}

func sumHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON payload
	var nums Numbers
	err := json.NewDecoder(r.Body).Decode(&nums)
	if err != nil {
		http.Error(w, "Failed to decode JSON payload", http.StatusBadRequest)
		return
	}

	// Calculate the sum
	sum := nums.Num1 + nums.Num2

	// Create the response
	response := SumResponse{Sum: sum}

	// Marshal the response to JSON
	respJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to marshal JSON response", http.StatusInternalServerError)
		return
	}

	// Set the response content type
	w.Header().Set("Content-Type", "application/json")

	// Write the response
	w.WriteHeader(http.StatusOK)
	w.Write(respJSON)
}

func main() {
	// Define the HTTP route
	http.HandleFunc("/sum", sumHandler)

	// Start the HTTP server
	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
