package handler

import (
	"Messenger/internal/core/domain"
	"Messenger/internal/core/services"

	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	svc services.MessengerService
}

func NewHTTPHandler(MessengerService services.MessengerService) *HTTPHandler {
	return &HTTPHandler{
		svc: MessengerService,
	}
}

func (h *HTTPHandler) SaveMessage(ctx *gin.Context) {
	var message domain.Message
	if err := ctx.ShouldBindJSON(&message); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err,
		})

		return
	}

	err := h.svc.SaveMessage(message)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "New message created successfully",
	})
}

func (h *HTTPHandler) ReadMessage(ctx *gin.Context) {
	id := ctx.Param("id")
	message, err := h.svc.ReadMessage(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, message)
}

func (h *HTTPHandler) ReadMessages(ctx *gin.Context) {

	messages, err := h.svc.ReadMessages()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, messages)
}
