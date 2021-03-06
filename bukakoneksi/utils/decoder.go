package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/harkce/beasiswakita/errors"
)

func Decode(r *http.Request, object interface{}) error {
	if r.Body == nil {
		return nil
	}

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.UnprocessableEntity
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	err = json.Unmarshal(bodyBytes, &object)
	if err != nil {
		return errors.UnprocessableEntity
	}

	return nil
}

func DecodeResponse(r *http.Response, object interface{}) error {
	if r.Body == nil {
		return nil
	}

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.UnprocessableEntity
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	err = json.Unmarshal(bodyBytes, &object)
	if err != nil {
		return errors.UnprocessableEntity
	}

	return nil
}
