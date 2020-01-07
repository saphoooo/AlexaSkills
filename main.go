package main

import (
	"fmt"
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
			switch slot := s.Request.Intent.Slots.SlotName.Name; slot {
			case "Foods":
				params.FoodName = s.Request.Intent.Slots.SlotName.Value
				resp, err := controllers.GetCookingInstructionIntent(params)
				if err != nil {
					panic(err)
				}
				fmt.Println(string(resp))
			case "DietTypes":
				params.DietTypes = s.Request.Intent.Slots.SlotName.Value
				resp, err := controllers.GetCookingInstructionIntent(params)
				if err != nil {
					panic(err)
				}
				fmt.Println(string(resp))
			default:
				log.Fatalln("unknow slot:", slot)
			}
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
