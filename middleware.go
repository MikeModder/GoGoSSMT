package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func debugMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debugf("%s - %s", r.RemoteAddr, r.RequestURI)

		_, err := sessionStore.Get(r, "ssmt-sess")
		if err != nil {
			log.Errorf("failed to get session for %s: %v", r.RemoteAddr, err)
			templates.ExecuteTemplate(w, "error.html", map[string]string{
				"Error": err.Error(),
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}
