package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"inventory/model"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"
const dbName = "inventory"
const empCollName = "employee"

var emp_collection *mongo.Collection

func init() {

	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect db
	client, err := mongo.Connect(context.Background(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	//CREATE COLLECTION INSTANCE
	emp_collection = client.Database(dbName).Collection(empCollName)

	fmt.Println("Employee Collection instance is ready...")
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee model.Employee
	_ = json.NewDecoder(r.Body).Decode(&employee)
	current := time.Now()
	employee.CreatedAt = current.Local()
	result, err := emp_collection.InsertOne(context.Background(), employee)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(result.InsertedID)
}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Control-Type", "application/json")
	cur, err := emp_collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var employees []model.Employee
	for cur.Next(context.Background()) {
		var employee model.Employee
		err := cur.Decode(&employee)
		if err != nil {
			log.Fatal(err)
		}
		employees = append(employees, employee)
	}
	defer cur.Close(context.Background())
	json.NewEncoder(w).Encode(employees)
}
func GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	empID, _ := primitive.ObjectIDFromHex(param["id"])
	filter := bson.M{"_id": empID}
	var employee model.Employee
	err := emp_collection.FindOne(context.Background(), filter).Decode(&employee)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(employee)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	eid := param["id"]
	empID, _ := primitive.ObjectIDFromHex(eid)
	filter := bson.M{"_id": empID}
	update := bson.M{"$set": bson.M{}}
	result, err := emp_collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(result.ModifiedCount)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	empid := param["id"]
	empID, _ := primitive.ObjectIDFromHex(empid)
	filter := bson.M{"_id": empID}
	delteCount, err := emp_collection.DeleteOne(context.Background(), filter, nil)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(delteCount)
}

func DeleteAllEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	deletedCount, err := emp_collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(deletedCount.DeletedCount)
}
