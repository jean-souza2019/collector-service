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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

type CreatedUserResponse struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func (u *User) validateUser() error {
	if u.Email == "" || u.Name == "" || u.Age <= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}

	if u.Email == "" {
		return errParamIsRequired("email", "string")
	}

	if u.Name == "" {
		return errParamIsRequired("name", "string")
	}

	if u.Age <= 0 {
		return errParamIsRequired("age", "number")
	}

	return nil
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var u User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		errorResponse(w, http.StatusBadRequest, "inválid parameters")
		return
	}

	if err := u.validateUser(); err != nil {
		errorMessage := fmt.Sprintf("validation error: %v", err.Error())

		fmt.Println(errorMessage)
		errorResponse(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	mongo.CheckConneciton()

	collection := mongo.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, u)
	if err != nil {
		log.Fatal(err)
	}

	objectId, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Fatal("Not accepted inserted id")
	}

	userCreated := CreatedUserResponse{
		Id:    objectId.Hex(),
		Email: u.Email,
		Name:  u.Name,
		Age:   u.Age,
	}

	fmt.Println("Success register user:", objectId.Hex())
	successResponse(w, "register-user", userCreated)
}

func FindUsersHandler(w http.ResponseWriter, r *http.Request) {
	mongo.CheckConneciton()

	var filter bson.D = bson.D{{}}
	var results []*CreatedUserResponse

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
		var elem CreatedUserResponse
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		userCreated := CreatedUserResponse{
			Id:    elem.Id,
			Email: elem.Email,
			Name:  elem.Name,
			Age:   elem.Age,
		}
		results = append(results, &userCreated)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	successResponse(w, "find-all-users", results)
}
