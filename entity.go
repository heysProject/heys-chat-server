package main

import "time"

type ChatReactionEntity struct {
	// Auto generated id
	ID string `json:"_id,omitempty" bson:"_id,omitempty"`
	// API Server user id (writter)
	ReactorId string `json:"writter_id" bson:"writter_id"`
	// Chat reactions
	Reaction string `json:"reaction" bson:"reaction"`

	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	RemovedAt time.Time `json:"removed_at,omitempty" bson:"removed_at,omitempty"`
}

type ChatEntity struct {
	// Auto generated id
	ID string `json:"_id,omitempty" bson:"_id,omitempty"`
	// API Server user id (writter)
	WritterId string `json:"writter_id" bson:"writter_id"`
	// MQTT Topic + '/' + Room ID
	Topic string `json:"topic" bson:"topic"`
	// Chat Message
	Message string `json:"message" bson:"message"`
	// Chat reactions
	Reaction []ChatReactionEntity `json:"reaction" bson:"reaction"`

	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	RemovedAt time.Time `json:"removed_at,omitempty" bson:"removed_at,omitempty"`
}

type ChatNoticeEntity struct {
	// Auto generated id
	ID string `json:"_id,omitempty" bson:"_id,omitempty"`
	// API Server user id (writter)
	WritterId string `json:"writter_id" bson:"writter_id"`
	// MQTT Topic + '/' + Room ID
	Topic string `json:"topic" bson:"topic"`
	// Chat Notice Title
	Title string `json:"title" bson:"title"`
	// Chat Notice Content
	Content string `json:"content" bson:"content"`
	// Chat reactions
	Reaction []ChatReactionEntity `json:"reaction" bson:"reaction"`

	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	RemovedAt time.Time `json:"removed_at,omitempty" bson:"removed_at,omitempty"`
}
