package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/gorilla/mux"
)

// Employee struct represents the employee data model
type Employee struct {
	ID         string    `json:"id,omitempty" bson:"_id,omitempty"`
	Name       string    `json:"name,omitempty" bson:"name,omitempty"`
	Age        int       `json:"age,omitempty" bson:"age,omitempty"`
	Department string    `json:"department,omitempty" bson:"department,omitempty"`
	CreatedAt  time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}

// Department struct represents the department data model
type Department struct {
	ID   string json:"id,omitempty" bson:"_id,omitempty"
	Name string json:"name,omitempty" bson:"name,omitempty"
}

var client *mongo.Client

// Connect to MongoDB
func connectDB() error {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	return nil
}

// Create an employee
func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee Employee
	_ = json.NewDecoder(r.Body).Decode(&employee)
	employee.CreatedAt = time.Now()
	collection := client.Database("employeesDB").Collection("employees")
	result, err := collection.InsertOne(context.Background(), employee)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte({"error": "Failed to create employee"}))
		return
	}
	json.NewEncoder(w).Encode(result.InsertedID)
}

// Get all employees
func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employees []Employee
	collection := client.Database("employeesDB").Collection("employees")
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte({"error": "Failed to get employees"}))
		return
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var employee Employee
		err := cur.Decode(&employee)
		if err != nil {
			log.Fatal(err)
		}
		employees = append(employees, employee)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(employees)
}

// Create a department
func createDepartment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var department Department
	_ = json.NewDecoder(r.Body).Decode(&department)
	collection := client.Database("employeesDB").Collection("departments")
	result, err := collection.InsertOne(context.Background(), department)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte({"error": "Failed to create department"}))
		return
	}
	json.NewEncoder(w).Encode(result.InsertedID)
}

// Get all departments
func getDepartments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var departments []Department
	collection := client.Database("employeesDB").Collection("departments")
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte({"error": "Failed to get departments"}))
		return
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var department Department
		err := cur.Decode(&department)
		if err != nil {
			log.Fatal(err)
		}
		departments = append(departments, department)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(departments)
}

// Main function
func main() {
	// Connect to MongoDB
	if err := connectDB(); err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Initialize router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/employees", createEmployee).Methods("POST")
	router.HandleFunc("/employees", getEmployees).Methods("GET")
	router.HandleFunc("/departments", createDepartment).Methods("POST")
	router.HandleFunc("/departments", getDepartments).Methods("GET")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", router))
}