package model

import (
	"errors"
	"net/http"

	"github.com/asaskevich/govalidator"
)

// InputValidation - an interface for all "input submission" structs used for
// deserialization.  We pass in the request so that we can potentially get the
// context by the request from our context manager
type InputValidation interface {
	Validate(r *http.Request) error
}

var (
	// ErrInvalidName - error when we have an invalid name
	SenderEmailReq     = errors.New("sender email empty or invalid")
	ReceiverEmailReq    = errors.New("receiver email empty or invalid")
)

func (t DraftMail) Validate(r *http.Request) error {
	// validate the email is not empty or missing
	if !govalidator.IsEmail(t.SenderEmail) {
		return SenderEmailReq
	}
	if !govalidator.IsEmail(t.RecipientEmail) {
		return ReceiverEmailReq
	}
	if govalidator.IsNull(t.Subject) {
		return ReceiverEmailReq
	}
	return nil
}
