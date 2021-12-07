package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"xanxit.com/helper"
	"xanxit.com/models"

	"github.com/gorilla/mux"
)

var collection = helper.ConnectDB()

func createBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var book models.Book
	// we decode our body request params
	_ = json.NewDecoder(req.Body).Decode(&book)
	// insert our book model.
	result, err := collection.InsertOne(context.TODO(), book)

	if err != nil {
		helper.GetError(err, res)
		return
	}

	json.NewEncoder(res).Encode(result)
}
func updateBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var book models.Book
	fmt.Println("hey sexy")
	filter := bson.M{"_id": id}
	_ = json.NewDecoder(req.Body).Decode(&book)
	update := bson.M{
		"$set": bson.D{
			{"title", book.Title},
		},
	}
	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&book)

	if err != nil {
		helper.GetError(err, res)
		return
	}
	book.ID = id
	json.NewEncoder(res).Encode(book)
}
func getBooks(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var books []models.Book
	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		helper.GetError(err, res)
		return
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var book models.Book
		err := cur.Decode(&book) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(res).Encode(books) // encode similar to serialize process.
}

func getBook(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")

	var book models.Book
	// we get params with mux.
	var params = mux.Vars(r)

	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&book)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(book)
}
func deleteBook(res http.ResponseWriter, req *http.Request) {
	// Set header
	res.Header().Set("Content-Type", "application/json")

	// get params
	var params = mux.Vars(req)

	// string to primitve.ObjectID
	id, err := primitive.ObjectIDFromHex(params["id"])

	// prepare filter.
	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		helper.GetError(err, res)
		return
	}

	json.NewEncoder(res).Encode(deleteResult)
}
