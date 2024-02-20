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

type Billing struct {
	Reference      string `json:"reference"`
	ContractId     int    `json:"contractId"`
	CustomerName   string `json:"customerName"`
	ExpirationDate string `json:"expirationDate"`
}

type CreatedBillingResponse struct {
	Id             string `json:"id"`
	Reference      string `json:"reference"`
	ContractId     int    `json:"contractId"`
	CustomerName   string `json:"customerName"`
	ExpirationDate string `json:"expirationDate"`
}

func (u *Billing) validateBilling() error {
	if u.Reference == "" || u.CustomerName == "" || u.ExpirationDate == "" || u.ContractId <= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}

	if u.Reference == "" {
		return errParamIsRequired("reference", "string")
	}

	if u.CustomerName == "" {
		return errParamIsRequired("customerName", "string")
	}

	if u.ExpirationDate == "" {
		return errParamIsRequired("expirationDate", "string")
	}

	if u.ExpirationDate != "" {
		layout := "2006-01-02"

		_, err := time.Parse(layout, u.ExpirationDate)

		if err != nil {
			return errParamIsRequired("expirationDate", "yyyy-mm-dd")
		}
	}

	if u.ContractId <= 0 {
		return errParamIsRequired("contractId", "number")
	}

	return nil
}

func CreateBillingHandler(w http.ResponseWriter, r *http.Request) {
	var u Billing

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		errorResponse(w, http.StatusBadRequest, "inválid parameters")
		return
	}

	if err := u.validateBilling(); err != nil {
		errorMessage := fmt.Sprintf("validation error: %v", err.Error())

		fmt.Println(errorMessage)
		errorResponse(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	mongo.CheckConneciton()

	collection := mongo.GetCollection("billings")
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

	billingCreated := CreatedBillingResponse{
		Id:             objectId.Hex(),
		Reference:      u.Reference,
		ExpirationDate: u.ExpirationDate,
		CustomerName:   u.CustomerName,
		ContractId:     u.ContractId,
	}

	fmt.Println("Success register billing:", objectId.Hex())
	successResponse(w, "register-billing", billingCreated)
}

func FindBillingsHandler(w http.ResponseWriter, r *http.Request) {
	mongo.CheckConneciton()

	var filter bson.D = bson.D{{}}
	var results []*Billing

	collection := mongo.GetCollection("billings")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var elem Billing
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

	err = json.NewEncoder(w).Encode(results)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
