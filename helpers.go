package main

import (
	"net/http"
	"strconv"

	"github.com/gomodule/redigo/redis"
	"github.com/saphoooo/AlexaSkills/spoonacular"
)

// jsonReply is a helper function to write a response with the appropriate header
func jsonReply(w http.ResponseWriter, reply []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(reply)
}

// offsetInc increments the offset by the desired value
func offsetInc(increment int, params *spoonacular.GetCookingParams) error {
	offset, err := strconv.Atoi(params.Offset)
	if err != nil {
		return err
	}
	offset += increment
	params.Offset = strconv.Itoa(offset)
	return nil
}

type cookingResponse struct {
	Params         *spoonacular.GetCookingParams
	Pool           *redis.Pool
	Key            string
	ResponseWriter http.ResponseWriter
}

type cookingSender interface {
	sendCookingResponse(cr *cookingResponse) error
}

// sendCooking returns the search results to the desired API
func sendCooking(cs cookingSender, cr *cookingResponse) error {
	err := cs.sendCookingResponse(cr)
	if err != nil {
		return err
	}
	return nil
}
