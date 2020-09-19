package model

import (
	"bytes"
	"encoding/json"
	"github.com/Jeffail/gabs/v2"
	"io/ioutil"
	"net/http"
)

func DecodeAndValidate(r *http.Request, v InputValidation) error {
	// json decode the payload - obviously this could be abstracted
	// to handle many content types
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	jsonParsed, _ := gabs.ParseJSON(body)
	payloadObject := jsonParsed.Path("payload").Bytes()
	if err := json.NewDecoder(bytes.NewReader(payloadObject)).Decode(v); err != nil {
		return err
	}

	defer r.Body.Close()
	// peform validation on the InputValidation implementation
	return v.Validate(r)
}