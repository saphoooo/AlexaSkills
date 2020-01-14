package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/saphoooo/AlexaSkills/google"
)

// actionsNewTextToSpeechResponse marshals a text in Google ActionsResponse format
func actionsNewTextToSpeechResponse(text string) ([]byte, error) {
	item := &google.Items{
		SimpleResponse: &google.SimpleResponse{
			TextToSpeech: text,
			DisplayText:  text,
		},
	}
	items := []*google.Items{item}
	action := &google.ActionsResponse{
		GPayload: &google.GPayload{
			Google: &google.Google{
				RichResponse: &google.RichResponse{
					Items: items,
				},
			},
		},
	}
	resp, err := json.Marshal(action)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type googleCookingResponse struct{}

// sendCookingResponse returns the search results to Actions on Google API
func (g googleCookingResponse) sendCookingResponse(cr *cookingResponse) error {
	cooking, err := getCookingInstructionIntent(cr.Params)
	if err != nil {
		return err
	}
	text, err := resultsToText(cooking)
	if err != nil {
		return err
	}
	resp, err := actionsNewTextToSpeechResponse(text)
	if err != nil {
		return err
	}
	err = keepCookingParams(cr.Pool, cr.Key, cr.Params)
	if err != nil {
		return err
	}
	jsonReply(cr.ResponseWriter, resp)
	return nil
}

// NewActionsVerifier ...
func NewActionsVerifier(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		secret := r.Header.Get("superSecret")
		if secret != os.Getenv("ZGUINGOUS_KITCHEN") {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	})
}
