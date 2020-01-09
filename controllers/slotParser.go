package controllers

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/saphoooo/AlexaSkills/views"
)

// SlotParser iterates over the slots included in the intend request
// see https://developer.amazon.com/fr-FR/docs/alexa/custom-skills/request-types-reference.html#intentrequest
func SlotParser(slot map[string]interface{}, params *views.GetCookingParams) error {
	for key := range slot {
		var newSlot views.Slot
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
