package repository

import (
	"Hexagonal-Architecture/internal/core/domain"
	"errors"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func (m *DB) CreateMessage(userID string, message domain.Message) error {
	message.ID = uuid.New().String()
	message = domain.Message{
		ID:     message.ID,
		UserID: userID,
		Body:   message.Body,
	}
	req := m.db.Create(&message)
	if req.RowsAffected == 0 {
		return fmt.Errorf("messages not saved: %v", req.Error)
	}
	return nil
}

func (m *DB) ReadMessage(id string) (*domain.Message, error) {
	message := &domain.Message{}
	req := m.db.First(&message, "id = ? ", id)
	if req.RowsAffected == 0 {
		return nil, errors.New("message not found")
	}
	return message, nil
}

func (m *DB) ReadMessages() ([]*domain.Message, error) {
	var messages []*domain.Message
	req := m.db.Find(&messages)
	if req.Error != nil {
		return nil, fmt.Errorf("messages not found: %v", req.Error)
	}
	return messages, nil
}

func (m *DB) UpdateMessage(id string, message domain.Message) error {
	req := m.db.Model(&message).Where("id = ?", id).Update(message)
	if req.RowsAffected == 0 {
		return errors.New("message not found")
	}
	return nil
}

func (m *DB) DeleteMessage(id string) error {
	message := &domain.Message{}
	req := m.db.Where("id = ?", id).Delete(&message)
	if req.RowsAffected == 0 {
		return errors.New("message not found")
	}
	return nil
}
