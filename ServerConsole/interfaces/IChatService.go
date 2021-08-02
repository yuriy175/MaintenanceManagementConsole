package interfaces

import (
	"../models"
)

// IChatService describes chat service interface
type IChatService interface {
	// Starts the service
	Start() 

	// GetChatNotes returns all chat notes from db
	GetChatNotes(equipName string) []models.ChatModel
}
