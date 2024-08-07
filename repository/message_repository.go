package repository

import (
	"github.com/kwa0x2/realtime-chat-backend/models"
	"gorm.io/gorm"
)

type MessageRepository struct {
	DB *gorm.DB
}

func (r *MessageRepository) CreateMessage(tx *gorm.DB, message *models.Message) (*models.Message, error) {
	db := r.DB
	if tx != nil {
		db = tx
	}

	if err := db.Create(&message).Error; err != nil {
		return nil, err
	}
	return message, nil
}

func (r *MessageRepository) GetPrivateConversation(senderId, receiverId string) ([]*models.Message, error) {
	var messages []*models.Message
	if err := r.DB.Where(
		"(message_sender_id = ? AND message_receiver_id = ?) OR (message_receiver_id = ? AND message_sender_id = ?)",
		senderId, receiverId, senderId, receiverId,
	).Find(&messages).Error; err != nil {
		return nil, err
	}

	return messages, nil
}
