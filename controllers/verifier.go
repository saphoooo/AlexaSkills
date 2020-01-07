package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/pkg/errors"

	"cooking.io/views"
)

// Verifier checks the request's accuracy
func Verifier(r *http.Request) (*views.AlexaRequest, error) {
	var s views.AlexaRequest
	a, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.WithMessage(err, "whoops an error occurs with the verifier")
	}
	err = json.Unmarshal(a, &s)
	if err != nil {
		return nil, errors.WithMessage(err, "unable to unmarshal the request body")
	}
	if s.Context.System.Application.ApplicationID != os.Getenv("ALEXA_SKILLID") {
		return nil, errors.New("applicationID mismatch")
	}
	return &s, nil

}
