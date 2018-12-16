package static

import (
	"go-hello/internal/services/mailsend"
	"html/template"
	"net/http"
)

//Contact will load the contact static page
func Contact(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/template/layouts/main.html", "web/template/contact.html"))
	if r.Method != http.MethodPost {
		tmpl.ExecuteTemplate(w, "main", nil)
		return
	}

	message := "From: \n" + r.FormValue("email") + "\n\nMessage:\n" + r.FormValue("message")
	messageSent := mailsend.Send(r.FormValue("subject"), message)

	tmpl.ExecuteTemplate(w, "main", messageSent)
}

//Home loads the index/home page
func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/template/layouts/main.html", "web/template/home.html"))
	tmpl.ExecuteTemplate(w, "main", nil)
}
