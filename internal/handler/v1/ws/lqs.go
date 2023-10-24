package ws

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	wsm "github.com/Live-Quiz-Project/Backend/internal/model/ws"
	lqsUtil "github.com/Live-Quiz-Project/Backend/internal/util/lqs"

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
	type request struct {
		QuizID string `json:"quizId"`
	}

	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lqsID := uuid.New().String()

	var code string
	codes := make([]string, 0)
	for _, s := range h.hub.LiveQuizSessions {
		if s.QuizID == req.QuizID {
			lqsID = s.ID
		}
		codes = append(codes, s.Code)
	}

	if _, ok := h.hub.LiveQuizSessions[lqsID]; !ok {
		code = lqsUtil.LQSCodeGenerator(codes)
		h.hub.LiveQuizSessions[lqsID] = &wsm.LiveQuizSession{
			ID:      lqsID,
			QuizID:  req.QuizID,
			Code:    code,
			Clients: make(map[string]*wsm.Client),
			Moderator: &wsm.Moderator{
				LiveQuizSessionID: lqsID,
				CurrentQuestion:   0,
				Status:            lqsUtil.Idle,
			},
		}
		c.JSON(http.StatusOK, struct {
			ID     string `json:"id"`
			QuizID string `json:"quizId"`
			Code   string `json:"code"`
		}{
			ID:     lqsID,
			QuizID: req.QuizID,
			Code:   code,
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
	lqsID := c.Param("id")
	if _, ok := h.hub.LiveQuizSessions[lqsID]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No such session exists"})
		return
	}

	for _, cl := range h.hub.LiveQuizSessions[lqsID].Clients {
		if err := cl.Conn.Close(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Can't close client connection"})
			return
		}
	}

	h.hub.Broadcast <- &wsm.Message{
		Content: wsm.Content{
			Type:    lqsUtil.EndLQS,
			Payload: "Session has ended.",
		},
		LiveQuizSessionID: lqsID,
		UserID:            "Host",
		IsHost:            true,
	}

	delete(h.hub.LiveQuizSessions, lqsID)
	c.JSON(http.StatusOK, gin.H{"message": "Successfully ended the session"})
}

func (h *WSHandler) GetLiveQuizSessions(c *gin.Context) {
	lqsID := c.Param("id")

	lqses := make([]struct {
		ID     string `json:"id"`
		QuizID string `json:"quizId"`
		Code   string `json:"code"`
	}, 0)

	if lqsID != "" {
		for _, s := range h.hub.LiveQuizSessions {
			if s.ID == lqsID {
				c.JSON(http.StatusOK, struct {
					ID     string `json:"id"`
					QuizID string `json:"quizId"`
					Code   string `json:"code"`
				}{
					ID:     s.ID,
					QuizID: s.QuizID,
					Code:   s.Code,
				})
				return
			}
		}
	}

	uid := c.Query("uid")
	code := c.Query("code")
	quizID := c.Query("quiz-id")

	if (uid != "" && code != "") || (uid != "" && quizID != "") || (code != "" && quizID != "") || (uid != "" && code != "" && quizID != "") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only one query parameter is allowed."})
		return
	}
	if uid != "" {
		for _, s := range h.hub.LiveQuizSessions {
			for _, c := range s.Clients {
				if uid == c.ID {
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
			}
		}
		c.JSON(http.StatusOK, lqses)
		return
	}
	if code != "" {
		for _, s := range h.hub.LiveQuizSessions {
			if code == s.Code {
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
		}
		c.JSON(http.StatusOK, lqses)
		return
	}
	if quizID != "" {
		for _, s := range h.hub.LiveQuizSessions {
			if quizID == s.QuizID {
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
		}

		c.JSON(http.StatusOK, lqses)
		return
	}

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
}

func (h *WSHandler) GetParticipants(c *gin.Context) {
	var p []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	lqsID := c.Param("id")
	if _, ok := h.hub.LiveQuizSessions[lqsID]; !ok {
		p = make([]struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		}, 0)
		c.JSON(http.StatusOK, p)
		return
	}

	for _, c := range h.hub.LiveQuizSessions[lqsID].Clients {
		if !c.IsHost {
			p = append(p, struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			}{
				ID:   c.ID,
				Name: c.Name,
			})
		}
	}
	c.JSON(http.StatusOK, p)
}

func (h *WSHandler) GetHost(c *gin.Context) {
	var p []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	lqsID := c.Param("id")

	if _, ok := h.hub.LiveQuizSessions[lqsID]; !ok {
		p = make([]struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		}, 0)
		c.JSON(http.StatusOK, p)
		return
	}

	for _, c := range h.hub.LiveQuizSessions[lqsID].Clients {
		if c.IsHost {
			p = append(p, struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			}{
				ID:   c.ID,
				Name: c.Name,
			})
			break
		}
	}
	c.JSON(http.StatusOK, p)
}

func (h *WSHandler) CheckLiveQuizSessionAvailability(c *gin.Context) {
	code := c.Param("code")

	for _, s := range h.hub.LiveQuizSessions {
		if s.Code == code {
			c.JSON(http.StatusOK, struct {
				ID              string `json:"id"`
				QuizID          string `json:"quizId"`
				Code            string `json:"code"`
				CurrentQuestion int    `json:"curQ"`
				Status          string `json:"status"`
			}{
				ID:              s.ID,
				QuizID:          s.QuizID,
				Code:            s.Code,
				CurrentQuestion: s.Moderator.CurrentQuestion,
				Status:          s.Moderator.Status,
			})
			return
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "No such session exists"})
}

func (h *WSHandler) JoinLiveQuizSession(c *gin.Context) {
	code := c.Param("code")
	uid := c.Query("uid")
	uname := c.Query("uname")
	isHost := c.Query("is-host")
	var lqsID string

	for _, s := range h.hub.LiveQuizSessions {
		if s.Code == code {
			lqsID = s.ID
		}
		for _, cl := range s.Clients {
			if cl.IsHost && isHost == "true" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Host already exists in this session"})
				return
			}
		}
	}

	if _, ok := h.hub.LiveQuizSessions[lqsID]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No such session exists"})
		return
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
			Type:    lqsUtil.JoinedLQS,
			Payload: nil,
		},
		LiveQuizSessionID: lqsID,
		UserID:            uid,
		IsHost:            isHost == "true",
	}

	go writeMessage(cl)
	h.readMessage(cl)
}

func writeMessage(c *wsm.Client) {
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

func (h *WSHandler) readMessage(c *wsm.Client) {
	defer func() {
		h.hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("Error occured: %v", err)
				break
			}
		}
		log.Printf("Message received: %s", m)

		var mstr wsm.Content
		if err := json.Unmarshal(m, &mstr); err != nil {
			log.Printf("Error occured: %v", err)
			break
		}

		switch mstr.Type {
		case lqsUtil.JoinedLQS:
			h.sendMessage(c, mstr)
		case lqsUtil.LeftLQS:
			h.sendMessage(c, mstr)
		case lqsUtil.StartLQS:
			h.startLiveQuizSession(c, mstr.Payload)
		case lqsUtil.DistQuestion:
			h.distributeQuestion(mstr.Payload)
		case lqsUtil.SetLQSStatus:
			h.setLiveQuizSessionStatus(mstr.Payload)
		default:
			h.sendMessage(c, mstr)
		}
	}
}

func (h *WSHandler) sendMessage(c *wsm.Client, ct wsm.Content) {
	msg := &wsm.Message{
		Content: wsm.Content{
			Type:    ct.Type,
			Payload: ct.Payload,
		},
		LiveQuizSessionID: c.LiveQuizSessionID,
		UserID:            c.ID,
		IsHost:            c.IsHost,
	}

	if msg.IsHost {
		h.hub.Broadcast <- msg
	} else {
		h.hub.Private <- msg
	}
}

func (h *WSHandler) startLiveQuizSession(c *wsm.Client, payload interface{}) {
	msg := &wsm.Message{
		Content: wsm.Content{
			Type:    lqsUtil.StartLQS,
			Payload: nil,
		},
		LiveQuizSessionID: c.LiveQuizSessionID,
		UserID:            c.ID,
		IsHost:            c.IsHost,
	}
	h.hub.Broadcast <- msg

	p, ok := payload.(map[string]interface{})
	if !ok {
		log.Printf("Error occured: %v", "Payload is not a map")
		return
	}
	lqsID := p["id"].(string)
	if _, ok := h.hub.LiveQuizSessions[lqsID]; !ok {
		return
	}

	h.hub.Moderate <- &wsm.Moderator{
		LiveQuizSessionID: lqsID,
		CurrentQuestion:   0,
		Status:            lqsUtil.Starting,
	}

	done := make(chan struct{})
	go h.countdown(5, lqsID, wsm.Moderator{
		LiveQuizSessionID: lqsID,
		CurrentQuestion:   0,
		Status:            lqsUtil.Starting,
	}, done)
	<-done

	h.distributeQuestion(payload)
}

func (h *WSHandler) distributeQuestion(payload interface{}) {
	p, ok := payload.(map[string]interface{})
	if !ok {
		log.Printf("Error occured: %v", "Payload is not a map")
		return
	}
	lqsID := p["id"].(string)
	if _, ok := h.hub.LiveQuizSessions[lqsID]; !ok {
		return
	}
	var hostID string
	for _, cl := range h.hub.LiveQuizSessions[lqsID].Clients {
		if cl.IsHost {
			hostID = cl.ID
			break
		}
	}

	h.hub.Moderate <- &wsm.Moderator{
		LiveQuizSessionID: lqsID,
		CurrentQuestion:   0,
		Status:            lqsUtil.Questioning,
	}

	h.hub.Broadcast <- &wsm.Message{
		Content: wsm.Content{
			Type:    lqsUtil.DistQuestion,
			Payload: nil,
		},
		LiveQuizSessionID: lqsID,
		UserID:            hostID,
		IsHost:            true,
	}
}

func (h *WSHandler) setLiveQuizSessionStatus(payload interface{}) {
	p, ok := payload.(map[string]interface{})
	if !ok {
		log.Printf("Error occured: %v", "Payload is not a map")
		return
	}
	lqsID := p["id"].(string)
	if _, ok := h.hub.LiveQuizSessions[lqsID]; !ok {
		return
	}

	h.hub.Moderate <- &wsm.Moderator{
		LiveQuizSessionID: lqsID,
		CurrentQuestion:   0,
		Status:            p["status"].(string),
	}
}

func (h *WSHandler) countdown(seconds int, lqsID string, mod wsm.Moderator, cd chan<- struct{}) {
	for i := seconds; i >= 0; i-- {
		h.hub.Broadcast <- &wsm.Message{
			Content: wsm.Content{
				Type: lqsUtil.Countdown,
				Payload: struct {
					LiveQuizSessionID string `json:"lqsId"`
					Timeleft          int    `json:"timeLeft"`
					CurrentQuestion   int    `json:"curQ"`
					Status            string `json:"status"`
				}{
					LiveQuizSessionID: lqsID,
					Timeleft:          i,
					CurrentQuestion:   mod.CurrentQuestion,
					Status:            mod.Status,
				},
			},
			LiveQuizSessionID: lqsID,
			UserID:            "Countdown",
			IsHost:            true,
		}
		if i > 0 {
			time.Sleep(time.Second)
		}
	}
	close(cd)
}
