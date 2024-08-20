package handlers

import (
	"net/http"
	"text/template"

	"github.com/LucasWiman90/SimpleWeb/models"
)

// ViewHandler is the handler for viewing individual articles
func ViewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := models.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

// EditHandler is the handler for editing individual articles
// If an article page doesn't exist, create a new page with that name
// that can be saved.
func EditHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := models.LoadPage(title)
	if err != nil {
		p = &models.Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

// SaveHandler is the handler for saving individual articles
func SaveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save"):]
	body := r.FormValue("body")
	p := &models.Page{Title: title, Body: []byte(body)}
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

// renderTemplate takes a html template file and renders it
func renderTemplate(w http.ResponseWriter, tmpl string, p *models.Page) {
	t, err := template.ParseFiles("./templates/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
