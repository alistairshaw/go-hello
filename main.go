package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"go-hello/internal/services/config"
	"go-hello/internal/services/mailsend"
)

//ToolPageData contains template variables for the tools page
type ToolPageData struct {
	PageTitle    string
	SelectedTool string
}

//SpellPageData contains template variables for the spells page
type SpellPageData struct {
	PageTitle     string
	SelectedSpell string
}

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if config.LogRoutesToConsole() {
			log.Println(r.URL.Path)
		}
		f(w, r)
	}
}

func spell(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := SpellPageData{
		PageTitle:     vars["spell"],
		SelectedSpell: vars["spell"]}
	tmpl := template.Must(template.ParseFiles("web/template/layouts/main.html", "web/template/spell.html"))
	tmpl.ExecuteTemplate(w, "main", data)
}

func contact(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/template/layouts/main.html", "web/template/contact.html"))
	if r.Method != http.MethodPost {
		tmpl.ExecuteTemplate(w, "main", nil)
		return
	}

	message := "From: \n" + r.FormValue("email") + "\n\nMessage:\n" + r.FormValue("message")
	messageSent := mailsend.Send(r.FormValue("subject"), message)

	tmpl.ExecuteTemplate(w, "main", messageSent)
}

func tool(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := ToolPageData{
		PageTitle:    vars["tool"],
		SelectedTool: vars["tool"]}
	tmpl := template.Must(template.ParseFiles("web/template/layouts/main.html", "web/template/tool.html"))
	tmpl.ExecuteTemplate(w, "main", data)
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/template/layouts/main.html", "web/template/home.html"))
	tmpl.ExecuteTemplate(w, "main", nil)
}

func main() {
	r := mux.NewRouter()

	// serve static files
	fs := http.FileServer(http.Dir("./web/"))
	r.PathPrefix("/static/").Handler(fs)

	// home
	r.HandleFunc("/", logging(home))

	// api endpoints
	r.HandleFunc("/tools/{tool}", logging(tool))
	r.HandleFunc("/spells/{spell}", logging(spell))
	r.HandleFunc("/contact", logging(contact))

	http.ListenAndServe(":80", r)
}
