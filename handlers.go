package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Serve up the index page
}

func testError(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "iserror", map[string]string{
		"Error": "this is a demo error, fear me!",
	})
}

func notImpl(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Route %s not implemented! Soon:tm:", r.URL.String())
}
