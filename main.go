package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Handle("/alexa", NewSkillsVerifier(http.HandlerFunc(alexaskills))).Methods("POST")
	r.Handle("/actions", NewActionsVerifier(http.HandlerFunc(actions)))
	log.Println("Start listening on :8000...")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		panic(err)
	}
}
