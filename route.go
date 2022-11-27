package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func route(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "go chat server"})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "chat server: pong!"})
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

		if publishChat(dto, roomId) {
			c.JSON(http.StatusOK, gin.H{"message": "message sent"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "message not sent"})
		}
	})

	router.Run()
}
