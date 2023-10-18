package ws

import (
	"encoding/json"
	"log"

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
	Username          string  `json:"uname"`
	IsHost            bool    `json:"isHost"`
	IsBroadcast       bool    `json:"isBroadcast"`
}

type Content struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

func (c *Client) WriteMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		m, ok := <-c.Message
		if !ok {
			return
		}

		c.Conn.WriteJSON(m)
	}
}

func (c *Client) ReadMessage(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("Error occured: %v", err)
			}
			break
		}

		var ct Content
		er := json.Unmarshal([]byte(m), &ct)
		if er != nil {
			msg := &Message{
				Content:           ct,
				LiveQuizSessionID: c.LiveQuizSessionID,
				UserID:            c.ID,
				Username:          c.Name,
				IsHost:            c.IsHost,
				IsBroadcast:       c.IsHost,
			}

			if msg.IsBroadcast {
				hub.Broadcast <- msg
			} else {
				hub.Private <- msg
			}
		}
	}
}
