package services

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"polymail/app/model"
	"polymail/app/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetSession(ctx context.Context) (*mongo.Client, error) {
	// Connect to our local mongo
	fmt.Println(os.Getenv("MONGODB_URL"))
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URL")))
	return client, err
}

type DbSession struct {
	mgo *mongo.Client
}

func DbClient(conn *mongo.Client) repository.DraftMailDB {
	return &DbSession{mgo: conn}
}

func (s *DbSession) CreateDraftMail(mail *model.DraftMail) error {
	mail.ID = primitive.NewObjectID()
	c := s.mgo.Database("polymail").Collection("draftemail")
	if _, err := c.InsertOne(context.Background(), mail); err != nil {
		return err
	}
	return nil
}

func (s *DbSession) GetDraftMailById(mailId string, u *model.DraftMail) error {
	objectId, err := primitive.ObjectIDFromHex(mailId)
	if err != nil {
		return fmt.Errorf("invalid id")
	}

	c := s.mgo.Database("polymail").Collection("draftemail")
	if err := c.FindOne(context.Background(), bson.M{"_id": objectId}).Decode(u); err != nil {
		return fmt.Errorf("email does not exist")
	}
	return nil
}

func (s *DbSession) UpdateDraftMail(mailId string, data *model.DraftMail) error {
	objectId, err := primitive.ObjectIDFromHex(mailId)
	if err != nil {
		return fmt.Errorf("invalid id")
	}

	c := s.mgo.Database("polymail").Collection("draftemail")
	change := bson.M{"$set": bson.M{
		"sender_email":    data.SenderEmail,
		"recipient_email": data.RecipientEmail,
		"subject":         data.Subject,
		"message":         data.Message,
		"updated_at":      time.Now(),
	}}

	if _, err := c.UpdateOne(context.Background(), bson.M{"_id": objectId}, change); err != nil {
		return err
	}

	if err := c.FindOne(context.Background(), bson.M{"_id": objectId}).Decode(data); err != nil {
		return err
	}
	return nil
}

func (s *DbSession) DeleteDraftMail(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid id")
	}

	if _, err := s.mgo.Database("polymail").Collection("draftemail").
		DeleteOne(context.Background(), bson.M{"_id": objectId}); err != nil {
		return err
	}
	return nil
}

func (s *DbSession) SendDraftMail(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid id")
	}

	var draftMail model.DraftMail
	if err := s.mgo.Database("polymail").Collection("draftemail").
		FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&draftMail); err != nil {
		return err
	}

	mail := email{
		SenderEmail:   draftMail.SenderEmail,
		ReceiverEmail: draftMail.RecipientEmail,
		Subject:       draftMail.Subject,
		Message:       draftMail.Message,
	}

	if err := mail.SendDraftEmail(); err != nil {
		return err
	}

	if err := s.CreateSentMail(&draftMail); err != nil {
		return err
	}

	if err := s.DeleteDraftMail(id); err != nil {
		return err
	}

	return nil
}
