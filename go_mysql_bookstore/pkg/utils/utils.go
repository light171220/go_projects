package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) error {
	if r.Body == nil {
		return errors.New("request body is empty")
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.New("failed to read request body: " + err.Error())
	}

	if len(body) == 0 {
		return errors.New("request body is empty")
	}

	err = json.Unmarshal(body, x)
	if err != nil {
		return errors.New("failed to parse JSON: " + err.Error())
	}
	
	return nil
}