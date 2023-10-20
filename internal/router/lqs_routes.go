package router

import (
	"github.com/Live-Quiz-Project/Backend/internal/handler/v1/ws"
	"github.com/gin-gonic/gin"
)

func LiveQuizSessionRoutes(r *gin.Engine, wsHandler *ws.WSHandler) {
	v1 := r.Group("/v1")

	wsr := v1.Group("/lqses")
	wsr.GET("/", wsHandler.GetLiveQuizSessions)
	wsr.GET("/:id", wsHandler.GetLiveQuizSessions)
	wsr.GET("/:id/host", wsHandler.GetHost)
	wsr.GET("/:id/participants", wsHandler.GetParticipants)
	wsr.POST("/", wsHandler.CreateLiveQuizSession)
	wsr.DELETE("/:id", wsHandler.EndLiveQuizSession)
	wsr.GET("/join/:code", wsHandler.JoinLiveQuizSession)

	resr := wsr.Group("/responses")
	resr.POST("/", ws.ResponseHandler)
}
