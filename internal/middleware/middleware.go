package middleware

import (
	"go-hello/internal/services/config"
	"log"
	"net/http"
)

//All will run all middleware needed
func All(f http.HandlerFunc) http.HandlerFunc {
	return logRoute(f)
}

//logRoute will log the requested route to the console _if_ enabled in config
func logRoute(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if config.LogRoutesToConsole() {
			log.Println(r.URL.Path)
		}
		f(w, r)
	}
}
