package asciiartweb

import (
	"fmt"
	"net/http"
	"text/template"

	asciiartweb "asciiartweb/Functions"
)

type Error struct {
	Status string
	Type   string
}

func IndexPage(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		if r.Method == http.MethodGet {
			RenderTemplate(w, "./templates/index.html", nil, http.StatusOK)
		} else {
			a := &Error{Status: "400", Type: "Bad Request"}
			RenderTemplate(w, "./templates/errorPage.html", a, http.StatusBadRequest)
		}
	default:
		a := &Error{Status: "404", Type: "Page Not Found"}
		RenderTemplate(w, "./templates/errorPage.html", a, http.StatusNotFound)
	}
}

func AsciiArtPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		if len(r.FormValue("text")) > 200 {
			a := &Error{Status: "400", Type: "Bad Request"}
			RenderTemplate(w, "./templates/errorPage.html", a, http.StatusBadRequest)
			return
		}
		result := asciiartweb.AsciiArt(r.FormValue("text"), r.FormValue("banner"))
		RenderTemplate(w, "./templates/index.html", result, http.StatusOK)
	} else {
		a := &Error{Status: "400", Type: "Bad Request"}
		RenderTemplate(w, "./templates/errorPage.html", a, http.StatusBadRequest)
	}
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data any, status int) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, "We're sorry, but something went wrong on our end. Please try again later.", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	err = t.Execute(w, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}
}
