package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jean-souza2019/collector-service/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Name  string `bson:"name"`
	Age   int    `bson:"age"`
	Email string `bson:"email"`
}

func UserHandler(w http.ResponseWriter, r *http.Request) {

	if mongo.Client == nil {
		log.Fatal("MongoDB client não foi inicializado")
	}
	newUser := User{
		Name:  "John Doe",
		Age:   30,
		Email: "johndoe@example.com",
	}

	// Insere o novo usuário na coleção 'users'
	collection := mongo.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, newUser)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Usuário inserido com sucesso:", result.InsertedID)
	w.Header().Set("Content-Type", "application/json")

	err2 := json.NewEncoder(w).Encode(result)

	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}
}

func FindUsersHandler(w http.ResponseWriter, r *http.Request) {
	if mongo.Client == nil {
		log.Fatal("MongoDB client não foi inicializado")
	}

	var filter bson.D = bson.D{{}}
	var results []*User

	collection := mongo.GetCollection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Encontrando os documentos
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	// Iterando sobre os documentos encontrados
	for cursor.Next(ctx) {
		var elem User
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")

	err2 := json.NewEncoder(w).Encode(results)

	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}
}
