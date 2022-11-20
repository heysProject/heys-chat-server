package main

import (
	"context"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var mqttClient mqtt.Client
var mongoClient *mongo.Client
var chatModel *mongo.Collection

func initialize() {
	connectMqtt()
	mongoClient = connectMongo()
	chatModel = mongoClient.Database("heys-chat").Collection("chat")
}

func main() {
	initialize()
	router := gin.New()
	route(router)

	defer func() {
		fmt.Println("Closing mongo connection")
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
}
