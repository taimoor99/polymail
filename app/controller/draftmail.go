package controller

import (
	"github.com/go-chi/chi"
	"net/http"
	"polymail/app/model"
	"polymail/app/repository"
	"polymail/utils"
	"time"
)

type draftmail struct {
	db repository.DraftMailDB
}

func DraftMailRepository(ub repository.DraftMailDB) repository.DraftMail {
	return &draftmail{db: ub}
}

func (m draftmail) CreateDraftMailHandler(w http.ResponseWriter, r *http.Request) {
	var data model.DraftMail

	if err := model.DecodeAndValidate(r, &data); err != nil {
		utils.WriteJsonErr(w, err.Error())
		return
	}

	data.CreatedAt = time.Now()
	if err := m.db.CreateDraftMail(&data); err != nil {
		utils.WriteJsonErr(w, err.Error())
		return
	}

	utils.WriteJsonRes(w, data)
	return
}

func (m draftmail) GetDraftMailHandler(w http.ResponseWriter, r *http.Request) {
	var data model.DraftMail

	mailID := chi.URLParam(r, "id")
	if mailID == "" {
		utils.WriteJsonErr(w, "mail id not found in params")
		return
	}
	if err := m.db.GetDraftMailById(mailID, &data); err != nil {
		utils.WriteJsonErr(w, err.Error())
		return
	}

	utils.WriteJsonRes(w, data)
}

func (m draftmail) DeleteDraftMailHandler(w http.ResponseWriter, r *http.Request) {
	mailID := chi.URLParam(r, "id")
	if mailID == "" {
		utils.WriteJsonErr(w, "mail id not found in params")
		return
	}

	if err := m.db.DeleteDraftMail(mailID); err != nil {
		utils.WriteJsonErr(w, err.Error())
		return
	}

	utils.WriteJsonRes(w, nil)
	return
}

func (m draftmail) UpdateDraftMailHandler(w http.ResponseWriter, r *http.Request) {
	mailID := chi.URLParam(r, "id")
	if mailID == "" {
		utils.WriteJsonErr(w, "mail id not found in params")
		return
	}

	var data model.DraftMail
	if err := model.DecodeAndValidate(r, &data); err != nil {
		utils.WriteJsonErr(w, err.Error())
		return
	}

	if err := m.db.UpdateDraftMail(mailID, &data); err != nil {
		utils.WriteJsonErr(w, err.Error())
		return
	}

	utils.WriteJsonRes(w, data)
	return
}

func (m draftmail) SendDraftMailHandler(w http.ResponseWriter, r *http.Request) {
	mailID := chi.URLParam(r, "id")
	if mailID == "" {
		utils.WriteJsonErr(w, "mail id not found in params")
		return
	}

	if err := m.db.SendDraftMail(mailID); err != nil {
		utils.WriteJsonErr(w, err.Error())
		return
	}

	utils.WriteJsonRes(w, nil)
	return
}
