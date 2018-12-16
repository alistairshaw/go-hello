package tool

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

//pageData contains template variables for the tools page
type pageData struct {
	PageTitle    string
	SelectedTool string
}

//View shows the view for a given tool
func View(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := pageData{
		PageTitle:    vars["tool"],
		SelectedTool: vars["tool"]}
	tmpl := template.Must(template.ParseFiles("web/template/layouts/main.html", "web/template/tool.html"))
	tmpl.ExecuteTemplate(w, "main", data)
}
