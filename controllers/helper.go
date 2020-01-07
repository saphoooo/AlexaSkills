package controllers

import (
	"net/http"
)

// JSONReply ...
func JSONReply(w http.ResponseWriter, reply []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(reply)
}
