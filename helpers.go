package main

import (
	"net/http"
	"strconv"

	"cooking.io/spoonacular"
	"github.com/gomodule/redigo/redis"
)

func jsonReply(w http.ResponseWriter, reply []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(reply)
}

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
