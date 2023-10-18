package router

import (
	v1 "github.com/Live-Quiz-Project/Backend/internal/handler/v1"
	"github.com/gin-gonic/gin"
)

func WebSocketRoutes(r *gin.Engine, wsHandler *v1.WSHandler) {
	wsr := r.Group("/lqses")
	wsr.GET("/", wsHandler.GetLiveQuizSessions)
	wsr.GET("/:id", wsHandler.GetLiveQuizSessions)
	wsr.GET("/:id/host", wsHandler.GetHost)
	wsr.GET("/:id/participants", wsHandler.GetParticipants)
	wsr.POST("/", wsHandler.CreateLiveQuizSession)
	wsr.DELETE("/", wsHandler.EndLiveQuizSession)

	wsr.GET("/join/:code", wsHandler.JoinLiveQuizSession)
	// wsr.GET("/start/:code", wsHandler.StartLiveQuizSession)
}
