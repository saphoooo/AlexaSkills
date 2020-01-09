package controllers

import (
	"encoding/json"

	"github.com/saphoooo/AlexaSkills/views"
)

// NewPlainTextResponse is an helper for building a plaintext response
func NewPlainTextResponse(text string) ([]byte, error) {
	r := &views.AlexaResponse{
		Version: "1.0",
		Response: &views.Response{
			ShouldEndSession: true,
			OutputSpeech: &views.OutputSpeech{
				Type: "PlainText",
				Text: text,
			},
		},
	}
	m, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	return m, nil
}
