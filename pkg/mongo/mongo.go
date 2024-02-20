package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jean-souza2019/collector-service/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Connect() {
	uri := fmt.Sprintf("mongodb://%v:%v@%v:%v/?authSource=%v", configs.EnvMongoUser, configs.EnvMongoPassword, configs.EnvMongoHost, configs.EnvMongoPort, configs.EnvMongoUser)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to MongoDB established.")
	Client = client
}

func GetCollection(collectionName string) *mongo.Collection {
	return Client.Database("default").Collection(collectionName)
}

func CheckConneciton() {
	if Client == nil {
		log.Fatal("MongoDB client não foi inicializado")
	}
}
