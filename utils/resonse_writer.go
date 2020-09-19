package utils

import (
	"encoding/json"
	"net/http"
	"polymail/app/model"
)

func WriteJsonRes(w http.ResponseWriter, body interface{}) {
	res := model.JsonResponse{Message: "success", Success:true, Body:body}
	// Marshal provided interface into JSON structure
	uj, err := json.Marshal(res)
	if err != nil {
		WriteJsonErr(w, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(uj)
	return
}

func WriteJsonErr(w http.ResponseWriter, message string) {
	res := model.JsonResponse{Message: message, Success: false, Body: map[string]interface{}{}}
	uj, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.Write(uj)
	return
}
