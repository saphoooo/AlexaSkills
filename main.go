package main

import (
	"log"
	"net/http"

	"cooking.io/controllers"
	"github.com/gorilla/mux"
)

func alexa(w http.ResponseWriter, r *http.Request) {
	s, err := controllers.Verifier(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("unauthorized"))
		log.Println(err)
		return
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
