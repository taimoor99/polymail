package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/Jeffail/gabs/v2"
	"github.com/go-chi/chi"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"polymail/app/model"
	"polymail/app/routes"
	"polymail/app/services"
	"testing"
	"time"
)

var objectId string

func Init() (*chi.Mux, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	goc, err := services.GetSession(ctx)
	if err != nil {
		return nil, err
	}
	userDB := services.DbClient(goc)
	ctrl := DraftMailRepository(userDB)
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

	req, err := http.NewRequest("POST", ts.URL+"/addmaildraft", bytes.NewBuffer(payload))
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

	jsonParsed, err := gabs.ParseJSON(respBody)

	if err != nil {
		t.Fatal(err)
	}
	objectId, _ = jsonParsed.Path("body._id").Data().(string)
}

func TestGetDraftMailHandler(t *testing.T) {
	r, err := Init()
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL+"/getmaildraft/"+objectId, nil)
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

	req, err := http.NewRequest("PUT", ts.URL+"/updatemaildraft/"+objectId, bytes.NewBuffer(payload))
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

func TestDeleteDraftMailHandler(t *testing.T) {
	r, err := Init()
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, err := http.NewRequest("DELETE", ts.URL+"/deletemaildraft/"+objectId, nil)
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

	req, err := http.NewRequest("PUT", ts.URL+"/senddraftemail/"+objectId, nil)
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
