package controllers

import (
	"encoding/json"

	"cooking.io/views"
	"github.com/pkg/errors"
)

// SlotParser ...
func SlotParser(slot map[string]interface{}, params *views.GetCookingParams) error {
	for key := range slot {
		var newSlot views.Slot
		s, err := json.Marshal(slot[key])
		if err != nil {
			return errors.WithMessage(err, "error unmarshaling "+key)
		}
		err = json.Unmarshal(s, &newSlot)
		if err != nil {
			return errors.WithMessage(err, "error marshaling "+key)
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
