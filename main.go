package main

import (
	"flag"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

//ToolPageData : Template variables for the tools page
type ToolPageData struct {
	PageTitle    string
	SelectedTool string
}

//SpellPageData : Template variables for the spells page
type SpellPageData struct {
	PageTitle     string
	SelectedSpell string
}

func main() {
	var dir string
	flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("./"))
	r.PathPrefix("/static/").Handler(fs)

	r.HandleFunc("/tools/{tool}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		data := ToolPageData{
			PageTitle:    vars["tool"],
			SelectedTool: vars["tool"]}
		tmpl := template.Must(template.ParseFiles("web/template/layouts/main.html", "web/template/tool.html"))
		tmpl.ExecuteTemplate(w, "main", data)
	})

	r.HandleFunc("/spells/{spell}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		data := SpellPageData{
			PageTitle:     vars["spell"],
			SelectedSpell: vars["spell"]}
		tmpl := template.Must(template.ParseFiles("web/template/layouts/main.html", "web/template/spell.html"))
		tmpl.ExecuteTemplate(w, "main", data)
	})

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("web/template/layouts/main.html", "web/template/home.html"))
		tmpl.ExecuteTemplate(w, "main", nil)
	})

	http.ListenAndServe(":80", r)
}
