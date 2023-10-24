package ws

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn              *websocket.Conn
	Message           chan *Message
	ID                string `json:"uid"`
	Name              string `json:"uname"`
	IsHost            bool   `json:"isHost"`
	LiveQuizSessionID string `json:"lqsId"`
}

type Message struct {
	Content           Content `json:"content"`
	LiveQuizSessionID string  `json:"lqsId"`
	UserID            string  `json:"uid"`
	IsHost            bool    `json:"isHost"`
}

type Content struct {
	Type    string `json:"type"`
	Payload any    `json:"payload"`
}
