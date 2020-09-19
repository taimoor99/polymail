package repository

import (
	"net/http"
	"polymail/app/model"
)

type DraftMail interface {
	CreateDraftMailHandler(w http.ResponseWriter, r *http.Request)
	GetDraftMailHandler(w http.ResponseWriter, r *http.Request)
	DeleteDraftMailHandler(w http.ResponseWriter, r *http.Request)
	UpdateDraftMailHandler(w http.ResponseWriter, r *http.Request)
	SendDraftMailHandler(w http.ResponseWriter, r *http.Request)
}

type DraftMailDB interface {
	CreateDraftMail(mail *model.DraftMail) error
	GetDraftMailById(mailId string, mail *model.DraftMail) error
	UpdateDraftMail(mailId string, mail *model.DraftMail) error
	DeleteDraftMail(id string) error
	SendDraftMail(id string) error
}
