package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"cooking.io/controllers"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
)

func alexa(w http.ResponseWriter, r *http.Request) {
	s, err := controllers.Verifier(r)
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
		resp, err := controllers.NewPlainTextResponse("Welcome to Zingou's cooking app")
		if err != nil {
			panic(err)
		}
		controllers.JSONReply(w, resp)
	case "IntentRequest":
		switch intent := s.Request.Intent.Name; intent {
		case "GetCookingIntent":
			params := controllers.NewGetCookingParams()
			err := controllers.SlotParser(s.Request.Intent.Slots, params)
			if err != nil {
				panic(err)
			}
			cooking, err := controllers.GetCookingInstructionIntent(params)
			if err != nil {
				panic(err)
			}
			text, err := controllers.ResultsToText(cooking)
			if err != nil {
				panic(err)
			}
			resp, err := controllers.NewPlainTextResponse(text)
			if err != nil {
				panic(err)
			}
			err = controllers.KeepCookingParams(pool, s.Session.User.UserID, params)
			if err != nil {
				panic(err)
			}
			controllers.JSONReply(w, resp)
		case "GetMoreRecipesIntent":
			params, err := controllers.GetCookingParams(pool, s.Session.User.UserID)
			if err != nil {
				if err.Error() == "ErrNoKey" {
					resp, err := controllers.NewPlainTextResponse("You can say alexa ask zguingou's kitchen I want to cook burger.")
					if err != nil {
						panic(err)
					}
					controllers.JSONReply(w, resp)
					return
				}
				panic(err)
			}
			offset, err := strconv.Atoi(params.Offset)
			if err != nil {
				panic(err)
			}
			offset += 3
			params.Offset = strconv.Itoa(offset)

			cooking, err := controllers.GetCookingInstructionIntent(params)
			if err != nil {
				panic(err)
			}
			text, err := controllers.ResultsToText(cooking)
			if err != nil {
				panic(err)
			}
			resp, err := controllers.NewPlainTextResponse(text)
			if err != nil {
				panic(err)
			}
			err = controllers.KeepCookingParams(pool, s.Session.User.UserID, params)
			if err != nil {
				panic(err)
			}
			controllers.JSONReply(w, resp)
		case "AMAZON.HelpIntent":
			resp, err := controllers.NewPlainTextResponse("You can say alexa ask zguingou's kitchen I want to cook burger.")
			if err != nil {
				panic(err)
			}
			controllers.JSONReply(w, resp)
		default:
			log.Fatalln("unknow intent:", intent)
		}
	case "SessionEndedRequest":
		log.Println("Session ended:", s.Request.Reason)
	default:
		log.Fatalln("unknow request type:", requestType)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/alexa", alexa).Methods("POST")
	log.Println("Start listening on :8000...")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		panic(err)
	}
}
