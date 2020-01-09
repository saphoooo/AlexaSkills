package controllers

import (
	"net/http"
)

// JSONReply is an helper to quickly set an "application/json" header
// and write the response back
func JSONReply(w http.ResponseWriter, reply []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(reply)
}
