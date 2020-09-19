package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	DraftMail struct {
		ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
		Subject        string             `bson:"subject" json:"subject"`
		SenderEmail    string             `bson:"sender_email" json:"sender_email"`
		RecipientEmail string             `bson:"recipient_email" json:"recipient_email"`
		Message        string             `bson:"message" json:"message"`
		CreatedAt      time.Time          `bson:"created_at" json:"created_at"`
		UpdatedAt      time.Time          `bson:"updated_at" json:"updated_at"`
	}
	JsonResponse struct {
		Message string      `json:"message"`
		Success bool        `json:"success"`
		Body    interface{} `json:"body"`
	}
)
