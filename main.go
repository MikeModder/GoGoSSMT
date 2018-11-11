package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

	log "github.com/sirupsen/logrus"
)

var (
	logLevel     = log.DebugLevel
	port         = 8080
	sessionStore = sessions.NewCookieStore([]byte("a"))
	templates    = template.Must(template.ParseGlob("templates/**.html"))
)

func init() {
	log.Infoln("GoGoSSMT - Loading config")
	// TODO: load config
	log.Infoln("GoGoSSMT - Starting")
	log.SetLevel(logLevel)
	// Any extra startup stuff
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", notImpl)
	r.HandleFunc("/threads", notImpl)
	r.HandleFunc("/threads/new", notImpl)
	r.HandleFunc("/threads/{id}", notImpl)
	r.HandleFunc("/threads/{id}/rate", notImpl)
	r.HandleFunc("/threads/{id}/reply", notImpl)
	r.HandleFunc("/user/{id}", notImpl)

	if logLevel == log.DebugLevel {
		// Debug stuff woot
		r.HandleFunc("/error", testError)
	}

	r.Use(debugMiddleware)

	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
