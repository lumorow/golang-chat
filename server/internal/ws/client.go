package ws

import "github.com/gorilla/websocket"

type Client struct {
	Com      *websocket.Conn
	Message  chan *Message
	ID       string `json:"id"`
	RoomID   string `json:"roomid"`
	Username string `json:"username"`
}

type Message struct {
	Content  string `json:"content"`
	Room     string `json:"room"`
	Username string `json:"username"`
}
