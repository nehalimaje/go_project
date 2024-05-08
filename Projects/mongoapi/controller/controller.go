package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongoapi/model"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"
const dbName = "netflix"
const collName = "watchlist"

var collection *mongo.Collection

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB connected successfully!!!")

	//create collection instance
	collection = client.Database(dbName).Collection(collName)

	//collection instance is ready
	fmt.Println("Collection instance is ready")
}

//MONGO DB HELPERS FILE

// INSERT DATA
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie is inserted successfull!!!", inserted.InsertedID)

}

// UPDATE 1 DATA
func updateOneMovie(movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieID)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count of movie is :", result.ModifiedCount)
}

// DELETE 1 RECORD
func deleteOneMovie(movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie got deleted with count : ", deleteCount)
}

// DELETE ALL MOVIE
func deleteAllMovie() int64 {

	deletedCount, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted movie count :", deletedCount.DeletedCount)
	return deletedCount.DeletedCount
}

// GET ALL COLLECTION FROM DB
func getAllMovie() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)

		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())
	return movies
}

// ACTUAL CONTROLLER FILE
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovie := getAllMovie()
	json.NewEncoder(w).Encode(allMovie)
}

func GetOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var movie model.Netflix
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&movie)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(movie)
}

func CreatMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	updateOneMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	deleteOneMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}
