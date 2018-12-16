package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"go-hello/internal/middleware"
	"go-hello/internal/spell"
	"go-hello/internal/static"
	"go-hello/internal/tool"
)

func main() {
	r := mux.NewRouter()

	// serve static files
	fs := http.FileServer(http.Dir("./web/"))
	r.PathPrefix("/static/").Handler(fs)

	// static pages
	r.HandleFunc("/", middleware.All(static.Home))
	r.HandleFunc("/contact", middleware.All(static.Contact))

	// data driven pages
	r.HandleFunc("/tools/{tool}", middleware.All(tool.View))
	r.HandleFunc("/spells/{spell}", middleware.All(spell.View))

	http.ListenAndServe(":80", r)
}
