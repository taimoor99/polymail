package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"polymail/app/model"
	"polymail/app/routes"
	"testing"
	"time"
)

var objId string = "5fede04292295831e39b1a00"

type DbSessionMock struct {
	mock.Mock
}

func (d *DbSessionMock) CreateDraftMail(mail *model.DraftMail) error {
	mail.ID = primitive.NewObjectID()
	fmt.Println(mail.ID)
	return nil
}

func (d *DbSessionMock) GetDraftMailById(mailId string, mail *model.DraftMail) error {
	objectId, _ := primitive.ObjectIDFromHex(mailId)
	*mail = model.DraftMail{
		SenderEmail:    "taimoorshaukat6@gmail.com",
		RecipientEmail: "taimoor.emallates@gmail.com",
		Subject:        "test",
		Message:        "test",
		ID:             objectId,
	}
	return nil
}

func (d *DbSessionMock) UpdateDraftMail(mailId string, mail *model.DraftMail) error {
	objectId, _ := primitive.ObjectIDFromHex(mailId)
	*mail = model.DraftMail{
		SenderEmail:    "taimoorshaukat6@gmail.com",
		RecipientEmail: "taimoor.emallates@gmail.com",
		Subject:        "test updated",
		Message:        "test updated",
		ID:             objectId,
	}
	return nil
}

func (d *DbSessionMock) DeleteDraftMail(id string) error {
	return nil
}

func (d *DbSessionMock) SendDraftMail(id string) error {
	return nil
}

func Init() (*chi.Mux, error) {
	m := new(DbSessionMock)
	ctrl := DraftMailRepository(m)
	r := routes.NewRouter(ctrl)
	return r, nil
}

func TestCreateDraftMailHandler(t *testing.T) {
	r, err := Init()
	if err != nil {
		t.Fatal(err)
	}

	body := struct {
		Payload model.DraftMail `json:"payload"`
	}{
		Payload: model.DraftMail{
			SenderEmail:    "taimoorshaukat6@gmail.com",
			RecipientEmail: "taimoor.emallates@gmail.com",
			Subject:        "test",
			Message:        "test",
		},
	}

	payload, _ := json.Marshal(body)
	ts := httptest.NewServer(r)
	defer ts.Close()

	req, err := http.NewRequest("POST", ts.URL+"/createmaildraft", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var response model.JsonResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		t.Fatal(err)
	}

	if !response.Success {
		t.Fatal("success should be true", response.Message)
	}

	jsonParsed, err := gabs.ParseJSON(respBody)

	if err != nil {
		t.Fatal(err)
	}
	objId, _ = jsonParsed.Path("body._id").Data().(string)
	if len(objId) < 1 {
		t.Fatal("object id must not be nil or empty")
	}
}

func TestGetDraftMailHandler(t *testing.T) {
	r, err := Init()
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL+"/getmaildraft/"+objId, nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var response model.JsonResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		t.Fatal(err)
	}

	if !response.Success {
		t.Fatal("success should be true")
	}

	id, _ := primitive.ObjectIDFromHex(objId)
	faketime, _ := time.Parse("2014-09-12T11:45:26.371Z", "0001-01-01T00:00:00Z")
	expectedResult := model.DraftMail{
		SenderEmail:    "taimoorshaukat6@gmail.com",
		RecipientEmail: "taimoor.emallates@gmail.com",
		Subject:        "test",
		Message:        "test",
		ID:             id,
		CreatedAt:      faketime,
		UpdatedAt:      faketime,
	}

	byteData, err := json.Marshal(response.Body)
	if err != nil {
		t.Fatal("body must not be nil or empty")
	}

	var responsebody model.DraftMail
	if err := json.Unmarshal(byteData, &responsebody); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expectedResult, responsebody)
}

func TestUpdateDraftMailHandler(t *testing.T) {
	r, err := Init()
	if err != nil {
		t.Fatal(err)
	}

	body := struct {
		Payload model.DraftMail `json:"payload"`
	}{
		Payload: model.DraftMail{
			SenderEmail:    "taimoorshaukat6@gmail.com",
			RecipientEmail: "taimoor.emallates@gmail.com",
			Subject:        "test updated",
			Message:        "test updated",
		},
	}

	payload, _ := json.Marshal(body)

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, err := http.NewRequest("PUT", ts.URL+"/updatemaildraft/"+objId, bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var response model.JsonResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		t.Fatal(err)
	}

	if !response.Success {
		t.Fatal("success should be true")
	}

	id, _ := primitive.ObjectIDFromHex(objId)
	faketime, _ := time.Parse("2014-09-12T11:45:26.371Z", "0001-01-01T00:00:00Z")
	expectedResult := model.DraftMail{
		SenderEmail:    "taimoorshaukat6@gmail.com",
		RecipientEmail: "taimoor.emallates@gmail.com",
		Subject:        "test updated",
		Message:        "test updated",
		ID: id,
		CreatedAt: faketime,
		UpdatedAt: faketime,
	}

	byteData, err := json.Marshal(response.Body)
	if err != nil {
		t.Fatal("body must not be nil or empty")
	}

	var responsebody model.DraftMail
	if err := json.Unmarshal(byteData, &responsebody); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expectedResult, responsebody)
}

func TestDeleteDraftMailHandler(t *testing.T) {
	r, err := Init()
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, err := http.NewRequest("DELETE", ts.URL+"/deletemaildraft/"+objId, nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var response model.JsonResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		t.Fatal(err)
	}

	if !response.Success {
		t.Fatal("success should be true")
	}
}

func TestSendDraftMailHandler(t *testing.T) {

	TestCreateDraftMailHandler(t)

	r, err := Init()
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, err := http.NewRequest("PUT", ts.URL+"/senddraftemail/"+objId, nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var response model.JsonResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		t.Fatal(err)
	}

	if !response.Success {
		t.Fatal("success should be true")
	}
}
