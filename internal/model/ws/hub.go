package ws

import (
	"log"

	lqsUtil "github.com/Live-Quiz-Project/Backend/internal/util/lqs"
)

type Hub struct {
	LiveQuizSessions map[string]*LiveQuizSession
	Register         chan *Client
	Unregister       chan *Client
	Broadcast        chan *Message
	Private          chan *Message
	Moderate         chan *Moderator
}

type LiveQuizSession struct {
	ID        string             `json:"id"`
	QuizID    string             `json:"quizId"`
	Code      string             `json:"code"`
	Clients   map[string]*Client `json:"clients"`
	Moderator *Moderator         `json:"mod"`
}

type Moderator struct {
	LiveQuizSessionID string `json:"lqsId"`
	QuestionsAmount   int    `json:"qAmount"`
	CurrentQuestion   int    `json:"curQ"`
	Status            string `json:"status"`
}

func NewHub() *Hub {
	return &Hub{
		LiveQuizSessions: make(map[string]*LiveQuizSession),
		Register:         make(chan *Client),
		Unregister:       make(chan *Client),
		Broadcast:        make(chan *Message, 5),
		Private:          make(chan *Message, 5),
		Moderate:         make(chan *Moderator),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case cl := <-h.Register:
			if _, ok := h.LiveQuizSessions[cl.LiveQuizSessionID]; ok {
				lqs := h.LiveQuizSessions[cl.LiveQuizSessionID]

				if _, ok := lqs.Clients[cl.ID]; !ok {
					lqs.Clients[cl.ID] = cl
				}
			}
		case cl := <-h.Unregister:
			if cl.IsHost {
				log.Println("Host has left the session.")
				h.Broadcast <- &Message{
					Content: Content{
						Type:    lqsUtil.EndLQS,
						Payload: "Host has left the session.",
					},
					LiveQuizSessionID: cl.LiveQuizSessionID,
					UserID:            cl.ID,
					IsHost:            cl.IsHost,
				}
				if _, ok := h.LiveQuizSessions[cl.LiveQuizSessionID]; ok {
					for _, c := range h.LiveQuizSessions[cl.LiveQuizSessionID].Clients {
						delete(h.LiveQuizSessions[c.LiveQuizSessionID].Clients, c.ID)
						close(c.Message)
						c.Conn.Close()
					}
				}
			} else {
				if _, ok := h.LiveQuizSessions[cl.LiveQuizSessionID]; ok {
					if _, ok := h.LiveQuizSessions[cl.LiveQuizSessionID].Clients[cl.ID]; ok {
						if len(h.LiveQuizSessions[cl.LiveQuizSessionID].Clients) != 0 {
							h.Broadcast <- &Message{
								Content: Content{
									Type:    lqsUtil.LeftLQS,
									Payload: nil,
								},
								LiveQuizSessionID: cl.LiveQuizSessionID,
								UserID:            cl.ID,
								IsHost:            cl.IsHost,
							}
						}
						delete(h.LiveQuizSessions[cl.LiveQuizSessionID].Clients, cl.ID)
						close(cl.Message)
						cl.Conn.Close()
					}
				}
			}
		case m := <-h.Broadcast:
			if _, ok := h.LiveQuizSessions[m.LiveQuizSessionID]; ok {
				for _, cl := range h.LiveQuizSessions[m.LiveQuizSessionID].Clients {
					if cl.ID != m.UserID {
						cl.Message <- m
					}
				}
			}
		case m := <-h.Private:
			if _, ok := h.LiveQuizSessions[m.LiveQuizSessionID]; ok {
				for _, cl := range h.LiveQuizSessions[m.LiveQuizSessionID].Clients {
					if cl.IsHost {
						cl.Message <- m
					}
				}
			}
		case m := <-h.Moderate:
			if _, ok := h.LiveQuizSessions[m.LiveQuizSessionID]; ok {
				h.LiveQuizSessions[m.LiveQuizSessionID].Moderator = m
			}
		}
	}
}
