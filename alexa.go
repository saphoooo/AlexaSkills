package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/pkg/errors"
	"github.com/saphoooo/AlexaSkills/alexa"
	"github.com/saphoooo/AlexaSkills/spoonacular"
)

type alexaCookingResponse struct{}

// sendCookingResponse returns the search results to Alexa Skills API
func (a alexaCookingResponse) sendCookingResponse(cr *cookingResponse) error {
	text, err := resultsToText(getCookingInstructionIntent(cr.Params))
	if err != nil {
		return err
	}
	resp, err := skillsNewPlainTextResponse(text)
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

// skillsNewPlainTextResponse marshals a text in Alexa SkillsResponse format
func skillsNewPlainTextResponse(text string) ([]byte, error) {
	r := &alexa.SkillsResponse{
		Version: "1.0",
		Response: &alexa.Response{
			ShouldEndSession: true,
			OutputSpeech: &alexa.OutputSpeech{
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

// skillsSlotParser iterates over Alexa response slots to capture their content
// see https://developer.amazon.com/fr-FR/docs/alexa/custom-skills/request-types-reference.html#intentrequest
func skillsSlotParser(slot map[string]interface{}, params *spoonacular.GetCookingParams) error {
	for key := range slot {
		var newSlot alexa.Slot
		s, err := json.Marshal(slot[key])
		if err != nil {
			return errors.WithMessage(err, "unable to marshal "+key)
		}
		err = json.Unmarshal(s, &newSlot)
		if err != nil {
			return errors.WithMessage(err, "unable to unmarshal "+key)
		}
		if newSlot.Resolutions != nil {
			switch key {
			case "Foods":
				params.FoodName = newSlot.Resolutions.ResolutionsPerAuthority[0].Values[0].Value.Name
			case "DietTypes":
				params.DietTypes = newSlot.Resolutions.ResolutionsPerAuthority[0].Values[0].Value.Name
			default:
				return errors.WithMessage(err, "unknow slot: "+key)
			}
		}
	}
	return nil
}

// NewSkillsVerifier makes the necessary checks to ensure that the request comes from Alexa Skills
// see https://developer.amazon.com/fr-FR/docs/alexa/custom-skills/security-testing-for-an-alexa-skill.html#22-skills-hosted-as-web-services-on-your-own-endpoint
func NewSkillsVerifier(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buffer, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		// robotClone duplicates the request to perform validation on is content
		// instead of draining the original request body
		robotClone := ioutil.NopCloser(bytes.NewBuffer(buffer))
		original := ioutil.NopCloser(bytes.NewBuffer(buffer))
		var v alexa.Verifier
		err = json.NewDecoder(robotClone).Decode(&v)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if v.Context.System.Application.ApplicationID != os.Getenv("ALEXA_SKILLID") {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		r.Body = original
		h.ServeHTTP(w, r)
	})
}
