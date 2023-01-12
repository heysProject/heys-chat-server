package main

/*
Topic 규칙

Private 채팅:
- /chat/{roomId<channel_user_chat_room_id>}
-- channel_user_chat_room_id: API Server - ChannelUser entity의 chatRoomId 필드

Channel(Group) 채팅:
- /group/{roomId<channel_id>}
-- channel_user_chat_room_id: API Server - Channel entity의 id 필드
*/

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
