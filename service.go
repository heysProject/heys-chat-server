package main

import (
	"context"
	"time"
)

func publishChat(dto ChatPublishDTO, roomId string) bool {
	// publish msg to broker
	topic := dto.Topic + "/" + roomId
	token := mqttClient.Publish(topic, 0 /* QoS Level */, false, dto.Message)

	if token.Wait() {
		chat := ChatEntity{
			Message:   dto.Message,
			WritterId: dto.WritterId,
			Topic:     dto.Topic,
			Reaction:  []ChatReactionEntity{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			RemovedAt: time.Now(),
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		chatModel.InsertOne(ctx, chat)

		return true
	}

	return false
}
