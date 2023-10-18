package ws

import (
	"net/http"

	wsm "github.com/Live-Quiz-Project/Backend/internal/model/ws"
	"github.com/Live-Quiz-Project/Backend/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type WSHandler struct {
	hub *wsm.Hub
}

func NewWSHandler(h *wsm.Hub) *WSHandler {
	return &WSHandler{
		hub: h,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(req *http.Request) bool {
		// origin := req.Header.Get("Origin")
		// return origin == "http://localhost:5173" || origin == "http://localhost:5174" || origin == "http://localhost:3000"
		return true
	},
}

func (h *WSHandler) CreateLiveQuizSession(c *gin.Context) {
	var req struct {
		QuizID string `json:"quizId"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lqsID := uuid.New().String()

	var lqsCode string
	codes := make([]string, 0)
	for _, s := range h.hub.LiveQuizSessions {
		if s.QuizID == req.QuizID {
			lqsID = s.ID
		}
		codes = append(codes, s.Code)
	}

	if _, ok := h.hub.LiveQuizSessions[lqsID]; !ok {
		lqsCode = util.LQSCodeGenerator(codes)
		h.hub.LiveQuizSessions[lqsID] = &wsm.LiveQuizSession{
			ID:      lqsID,
			QuizID:  req.QuizID,
			Code:    lqsCode,
			Clients: make(map[string]*wsm.Client),
		}
		c.JSON(http.StatusOK, struct {
			ID     string `json:"id"`
			QuizID string `json:"quizId"`
			Code   string `json:"code"`
		}{
			ID:     lqsID,
			QuizID: req.QuizID,
			Code:   lqsCode,
		})
		return
	}

	c.JSON(http.StatusOK, struct {
		ID     string `json:"id"`
		QuizID string `json:"quizId"`
		Code   string `json:"code"`
	}{
		ID:     h.hub.LiveQuizSessions[lqsID].ID,
		QuizID: h.hub.LiveQuizSessions[lqsID].QuizID,
		Code:   h.hub.LiveQuizSessions[lqsID].Code,
	})
}

func (h *WSHandler) EndLiveQuizSession(c *gin.Context) {
	var req struct {
		LiveQuizSessionID string `json:"lqsId"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, ok := h.hub.LiveQuizSessions[req.LiveQuizSessionID]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No such session exists"})
		return
	}

	for _, cl := range h.hub.LiveQuizSessions[req.LiveQuizSessionID].Clients {
		if err := cl.Conn.Close(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Can't close client connection"})
			return
		}
	}

	delete(h.hub.LiveQuizSessions, req.LiveQuizSessionID)
	c.JSON(http.StatusOK, gin.H{"message": "Successfully ended the session"})
}

func (h *WSHandler) GetLiveQuizSessions(c *gin.Context) {
	id := c.Param("id")

	// if uid == "" {
	// 	c.JSON(http.StatusOK, []struct {
	// 		ID     string `json:"id"`
	// 		QuizID string `json:"quizId"`
	// 		Code   string `json:"code"`
	// 	}{})
	// }
	lqses := make([]struct {
		ID     string `json:"id"`
		QuizID string `json:"quizId"`
		Code   string `json:"code"`
	}, 0)
	if id == "" {
		for _, s := range h.hub.LiveQuizSessions {
			lqses = append(lqses, struct {
				ID     string `json:"id"`
				QuizID string `json:"quizId"`
				Code   string `json:"code"`
			}{
				ID:     s.ID,
				QuizID: s.QuizID,
				Code:   s.Code,
			})
		}
		c.JSON(http.StatusOK, lqses)
		return
	}

	// for _, s := range h.hub.LiveQuizSessions {
	// 	for _, c := range s.Clients {
	// 		if c.ID == id && c.IsHost {
	// 			lqses = append(lqses, struct {
	// 				ID     string `json:"id"`
	// 				QuizID string `json:"quizId"`
	// 				Code   string `json:"code"`
	// 			}{
	// 				ID:     s.ID,
	// 				QuizID: s.QuizID,
	// 				Code:   s.Code,
	// 			})
	// 		}
	// 	}
	// }
	// c.JSON(http.StatusOK, lqses)
}

func (h *WSHandler) GetParticipants(c *gin.Context) {
	var p []struct {
		ID       string `json:"id"`
		Username string `json:"uname"`
	}

	lqsID := c.Param("lqs-id")
	if _, ok := h.hub.LiveQuizSessions[lqsID]; !ok {
		p = make([]struct {
			ID       string `json:"id"`
			Username string `json:"uname"`
		}, 0)
		c.JSON(http.StatusOK, p)
		return
	}

	for _, c := range h.hub.LiveQuizSessions[lqsID].Clients {
		if !c.IsHost {
			p = append(p, struct {
				ID       string `json:"id"`
				Username string `json:"uname"`
			}{
				ID:       c.ID,
				Username: c.Name,
			})
		}
	}
	c.JSON(http.StatusOK, p)
}

func (h *WSHandler) GetHost(c *gin.Context) {
	var p []struct {
		ID       string `json:"id"`
		Username string `json:"uname"`
	}

	lqsID := c.Param("lqs-id")
	if _, ok := h.hub.LiveQuizSessions[lqsID]; !ok {
		p = make([]struct {
			ID       string `json:"id"`
			Username string `json:"uname"`
		}, 0)
		c.JSON(http.StatusOK, p)
		return
	}

	for _, c := range h.hub.LiveQuizSessions[lqsID].Clients {
		if c.IsHost {
			p = append(p, struct {
				ID       string `json:"id"`
				Username string `json:"uname"`
			}{
				ID:       c.ID,
				Username: c.Name,
			})
			break
		}
	}
	c.JSON(http.StatusOK, p)
}

func (h *WSHandler) JoinLiveQuizSession(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lqsID := c.Param("lqs-id")
	if _, ok := h.hub.LiveQuizSessions[lqsID]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No such session exists"})
		return
	}

	uid := c.Query("uid")
	uname := c.Query("uname")
	isHost := c.Query("is-host")

	cl := &wsm.Client{
		Conn:              conn,
		Message:           make(chan *wsm.Message, 10),
		ID:                uid,
		Name:              uname,
		IsHost:            isHost == "true",
		LiveQuizSessionID: lqsID,
	}

	h.hub.Register <- cl
	h.hub.Broadcast <- &wsm.Message{
		Content: wsm.Content{
			Type:    "joined-lqs",
			Payload: nil,
		},
		LiveQuizSessionID: lqsID,
		UserID:            uid,
		Username:          uname,
		IsHost:            isHost == "true",
		IsBroadcast:       true,
	}

	go cl.WriteMessage()
	cl.ReadMessage(h.hub)
}
