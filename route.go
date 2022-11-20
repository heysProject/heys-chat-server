package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type ChatPublishDTO struct {
	Topic   string `json:"topic"`
	Message string `json:"message"`
	UserId  string `json:"userId"`
}

type ChatMessage struct {
	ID       string `json:"_id,omitempty" bson:"_id,omitempty"`
	FullText string `json:"full_text,omitempty" bson:"full_text,omitempty"`
	UserId   string `json:"user_id,omitempty" bson:"user_id,omitempty"`
}

func route(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		// c.HTML(200, "index.html", nil)
		c.JSON(http.StatusOK, gin.H{"message": "go chat server"})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong!"})
	})

	router.GET("/chats", func(c *gin.Context) {
		var chatMessages []bson.M

		cursor, err := chatModel.Find(context.TODO(), bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		if err = cursor.All(context.TODO(), &chatMessages); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, chatMessages)
	})

	router.POST("/publish/:roomId", func(c *gin.Context) {
		var dto ChatPublishDTO
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		roomId := c.Param("roomId")

		fmt.Println("Topic: ", dto.Topic+"/"+roomId)

		// publish msg to broker
		token := mqttClient.Publish(dto.Topic+"/"+roomId, 0 /* QoS Level */, false, dto.Message)

		if token.Wait() {
			chat := ChatMessage{
				FullText: dto.Message,
				UserId:   dto.UserId,
			}

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			chatModel.InsertOne(ctx, bson.D{
				{Key: "message", Value: chat.FullText},
				{Key: "userId", Value: chat.UserId},
			})

			c.JSON(http.StatusOK, gin.H{"message": "success"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "failed"})
	})

	router.Run()
}
