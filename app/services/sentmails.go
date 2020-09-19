package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"polymail/app/model"
)

func (s *DbSession) CreateSentMail(mail *model.DraftMail) error {
	mail.ID = primitive.NewObjectID()
	c := s.mgo.Database("polymail").Collection("sentemail")
	if _, err := c.InsertOne(context.Background(), mail); err != nil {
		return err
	}
	return nil
}