package spell

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

//pageData contains template variables for the spells page
type pageData struct {
	PageTitle     string
	SelectedSpell string
}

//View shows the view spell page
func View(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := pageData{
		PageTitle:     vars["spell"],
		SelectedSpell: vars["spell"]}
	tmpl := template.Must(template.ParseFiles("web/template/layouts/main.html", "web/template/spell.html"))
	tmpl.ExecuteTemplate(w, "main", data)
}
