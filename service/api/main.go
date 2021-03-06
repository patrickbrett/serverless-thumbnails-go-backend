package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/apex/gateway"
)

// ContentType contains the Content-Type header sent on all responses.
const ContentType = "application/json; charset=utf8"

// MessageResponse models a simple message responses.
type MessageResponse struct {
	Message string `json:"message"`
}

// WelcomeMessageResponse is the response returned by the / endpoint.
var WelcomeMessageResponse = MessageResponse{"Welcome to the example API!"}

// AltWelcomeMessageResponse is the response returned by the /message endpoint.
var AltWelcomeMessageResponse = MessageResponse{"Here's another message!"}

// RootHandler is a http.HandlerFunc for the / endpoint.
func RootHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(WelcomeMessageResponse)
}

// MessageHandler is a http.HandlerFunc for the /message endpoint.
func MessageHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(AltWelcomeMessageResponse)
}

// RegisterRoutes registers the API's routes.
func RegisterRoutes() {
	http.Handle("/", h(RootHandler))
	http.Handle("/message", h(MessageHandler))
}

// h wraps a http.HandlerFunc and adds common headers.
func h(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", ContentType)
		next.ServeHTTP(w, r)
	})
}

func main() {
	RegisterRoutes()
	log.Fatal(gateway.ListenAndServe(":3000", nil))
}
