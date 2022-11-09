package constants

import (
	"time"

	"github.com/gorilla/websocket"
)

// Hub and Connection Middle Relationship Object
type Client struct {
	HubP     *Hub
	ConnP    *websocket.Conn
	SendChan chan []byte // Buffered channel of outbound messages.
}

type Subscription struct {
	ClientP *Client
	RoomId  string
}

type Message struct {
	Value  []byte
	RoomId string
}

// manage connections, broadcast messages
type Hub struct {
	// Registered clients.
	Rooms map[string]map[*Client]bool // chat room roomId(key):value(Client) 로 매핑

	// Inbound messages from the clients.
	Broadcast chan Message

	// Register requests from the clients.
	Register chan Subscription

	// Unregister requests from clients.
	Unregister chan Subscription
}

const (
	// Time allowed to write a message to the peer.
	WriteWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	PongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	PingPeriod = (PongWait * 9) / 10

	// Maximum message size allowed from peer.
	MaxMessageSize = 512
)

var (
	Newline = []byte{'\n'}
	Space   = []byte{' '}
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
