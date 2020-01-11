package main

import (
	"encoding/json"

	"cooking.io/google"
)

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
