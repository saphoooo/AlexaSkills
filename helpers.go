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

type sender interface {
	sendCookingResponse(c *cookingResponse) error
}

func sendCooking(s sender, cr *cookingResponse) error {
	err := s.sendCookingResponse(cr)
	if err != nil {
		return err
	}
	return nil
}
