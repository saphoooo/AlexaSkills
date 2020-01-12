package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/saphoooo/AlexaSkills/google"

	"github.com/gomodule/redigo/redis"
)

// This route capture messages from Alexa Skills
func alexaskills(w http.ResponseWriter, r *http.Request) {
	s, err := skillsVerifier(r)
	if err != nil {
		status := http.StatusUnauthorized
		w.WriteHeader(status)
		w.Write([]byte(http.StatusText(status)))
		log.Println(err)
		return
	}

	pool := &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}

	switch requestType := s.Request.Type; requestType {
	case "LaunchRequest":
		resp, err := skillsNewPlainTextResponse("Welcome to Zingou's cooking app")
		if err != nil {
			panic(err)
		}
		jsonReply(w, resp)
	case "IntentRequest":
		switch intent := s.Request.Intent.Name; intent {
		case "GetCookingIntent":
			params := newGetCookingParams()
			err := skillsSlotParser(s.Request.Intent.Slots, params)
			if err != nil {
				panic(err)
			}
			a := alexaCookingResponse{}
			c := &cookingResponse{
				Params:         params,
				Pool:           pool,
				Key:            s.Session.User.UserID,
				ResponseWriter: w,
			}

			err = sendCooking(a, c)
			if err != nil {
				panic(err)
			}
		case "GetMoreRecipesIntent":
			params, err := getCookingParams(pool, s.Session.User.UserID)
			if err != nil {
				if err.Error() == "ErrNoKey" {
					resp, err := skillsNewPlainTextResponse("You can say alexa ask zguingou's kitchen I want to cook burger.")
					if err != nil {
						panic(err)
					}
					jsonReply(w, resp)
					return
				}
				panic(err)
			}

			err = offsetInc(3, params)
			if err != nil {
				panic(err)
			}

			a := alexaCookingResponse{}
			c := &cookingResponse{
				Params:         params,
				Pool:           pool,
				Key:            s.Session.User.UserID,
				ResponseWriter: w,
			}

			err = sendCooking(a, c)
			if err != nil {
				panic(err)
			}

		case "AMAZON.HelpIntent":
			resp, err := skillsNewPlainTextResponse("You can say alexa ask zguingou's kitchen I want to cook burger.")
			if err != nil {
				panic(err)
			}
			jsonReply(w, resp)
		default:
			log.Fatalln("unknow intent:", intent)
		}
	case "SessionEndedRequest":
		log.Println("Session ended:", s.Request.Reason)
	default:
		log.Fatalln("unknow request type:", requestType)
	}
}

// This route capture messages from Actions on Google
func actions(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("superSecret") != os.Getenv("ZGUINGOUS_KITCHEN") {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
		return
	}

	pool := &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}

	var g google.AcionsRequest
	err := json.NewDecoder(r.Body).Decode(&g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	switch intent := g.QueryResult.Action; intent {
	case "input.cooking":
		params := newGetCookingParams()
		params.DietTypes = g.QueryResult.Parameters.DietTypes
		params.FoodName = g.QueryResult.Parameters.Foods
		ggl := googleCookingResponse{}
		c := &cookingResponse{
			Params:         params,
			Pool:           pool,
			Key:            g.OriginalDetectIntentRequest.Payload.Conversation.ConversationID,
			ResponseWriter: w,
		}

		err = sendCooking(ggl, c)
		if err != nil {
			panic(err)
		}
	case "input.more":
		params, err := getCookingParams(pool, g.OriginalDetectIntentRequest.Payload.Conversation.ConversationID)
		if err != nil {
			if err.Error() == "ErrNoKey" {
				resp, err := actionsNewTextToSpeechResponse("You can say I want to cook burger.")
				if err != nil {
					panic(err)
				}
				jsonReply(w, resp)
				return
			}
			panic(err)
		}

		err = offsetInc(3, params)
		if err != nil {
			panic(err)
		}

		ggl := googleCookingResponse{}
		c := &cookingResponse{
			Params:         params,
			Pool:           pool,
			Key:            g.OriginalDetectIntentRequest.Payload.Conversation.ConversationID,
			ResponseWriter: w,
		}

		err = sendCooking(ggl, c)
		if err != nil {
			panic(err)
		}
	case "input.help":
		resp, err := actionsNewTextToSpeechResponse("You can say I want to cook burger.")
		if err != nil {
			panic(err)
		}
		jsonReply(w, resp)
	default:
	}
}
