package httpctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/himmel520/practice2024/internal/usecase"
	"github.com/sirupsen/logrus"
)

type (
	Handler struct {
		uc       *usecase.Usecase
		upgrader *websocket.Upgrader
		log      *logrus.Logger
	}
)

func New(uc *usecase.Usecase, log *logrus.Logger) *Handler {
	return &Handler{uc: uc, upgrader: &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}, log: log}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		v1.GET("/ws", h.webSocketConn)
	}

	return r
}

func (h *Handler) sendError(ws *websocket.Conn, message string) {
	if err := ws.WriteJSON(respErr{message}); err != nil {
		h.log.Error("error sending error message: ", err)
	}
}
