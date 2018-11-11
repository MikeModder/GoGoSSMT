package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	log "github.com/sirupsen/logrus"
)

var port = 8080

func init() {
	log.Infoln("GoGoSSMT - Loading config")
	// TODO: load config
	log.Infoln("GoGoSSMT - Starting")
	// Any extra startup stuff
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", notImpl)

	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
