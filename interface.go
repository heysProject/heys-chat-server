package main

import "time"

/*
Topic 규칙

Private 채팅:
- /chat/{roomId<channel_user_chat_room_id>}
-- channel_user_chat_room_id: API Server - ChannelUser entity의 chatRoomId 필드

Channel(Group) 채팅:
- /group/{roomId<channel_id>}
-- channel_user_chat_room_id: API Server - Channel entity의 id 필드
*/

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

type ChatPublishDTO struct {
	// Chat room id not included
	Topic     string `json:"topic"`
	Message   string `json:"message"`
	WritterId string `json:"writter_id"`
}

type ChatModifyDTO struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

type ChatModifyReactionDTO struct {
	Id       string `json:"id"`
	Reaction string `json:"reaction"` // reaction + 1 or -1
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

type ChatNoticePublishDTO struct {
	// Chat room id not included
	Topic     string `json:"topic"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	WritterId string `json:"writter_id"`
}

type ChatNoticeModifyDTO struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
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
