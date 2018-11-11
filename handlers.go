package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {

}

func notImpl(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Route %s not implemented! Soon:tm:", r.URL.String())
}
