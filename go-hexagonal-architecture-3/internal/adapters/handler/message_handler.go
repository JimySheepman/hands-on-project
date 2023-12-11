package handler

import (
	"Hexagonal-Architecture/internal/adapters/repository"
	"Hexagonal-Architecture/internal/core/domain"
	"Hexagonal-Architecture/internal/core/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	svc services.MessengerService
}

func NewMessageHandler(MessengerService services.MessengerService) *MessageHandler {
	return &MessageHandler{
		svc: MessengerService,
	}
}

func (h *MessageHandler) CreateMessage(ctx *gin.Context) {
	apiCfg, err := repository.LoadAPIConfig()
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	// Validate token
	userID, err := ValidateToken(ctx.Request.Header.Get("Authorization"), apiCfg.JWTSecret)
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	var message domain.Message
	message.UserID = userID

	if err := ctx.ShouldBindJSON(&message); err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	err = h.svc.CreateMessage(userID, message)
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "message created successfully",
	})
}

func (h *MessageHandler) ReadMessage(ctx *gin.Context) {
	id := ctx.Param("id")
	message, err := h.svc.ReadMessage(id)

	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, message)
}

func (h *MessageHandler) ReadMessages(ctx *gin.Context) {

	messages, err := h.svc.ReadMessages()

	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, messages)
}

func (h *MessageHandler) UpdateMessage(ctx *gin.Context) {
	apiCfg, err := repository.LoadAPIConfig()
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	// Validate token
	userID, err := ValidateToken(ctx.Request.Header.Get("Authorization"), apiCfg.JWTSecret)
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	// check if userID match with message.UserID
	id := ctx.Param("id")
	msg, err := h.svc.ReadMessage(id)
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}
	if msg.UserID != userID {
		HandleError(ctx, http.StatusBadRequest, fmt.Errorf("you are not authorized to update this message"))
		return
	}

	var message domain.Message
	if err := ctx.ShouldBindJSON(&message); err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	err = h.svc.UpdateMessage(id, message)
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Message updated successfully",
	})
}

func (h *MessageHandler) DeleteMessage(ctx *gin.Context) {
	apiCfg, err := repository.LoadAPIConfig()
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	userID, err := ValidateToken(ctx.Request.Header.Get("Authorization"), apiCfg.JWTSecret)
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	// check if userID match with message.UserID
	id := ctx.Param("id")
	message, err := h.svc.ReadMessage(id)
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}
	if message.UserID != userID {
		HandleError(ctx, http.StatusBadRequest, fmt.Errorf("you are not authorized to delete this message"))
		return
	}

	err = h.svc.DeleteMessage(id)
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Message deleted successfully",
	})
}
