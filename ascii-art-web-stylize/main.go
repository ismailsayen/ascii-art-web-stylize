package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	asciiartweb "asciiartweb/handlers"
)

const port string = ":8080"

type Error struct {
	Status string
	Type   string
}

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Enter 1 arg")
		return
	}
	fs := http.FileServer(http.Dir("assets"))

	http.HandleFunc("/assets/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/assets/" {
			a := Error{Status: "404", Type: "Page Not Found"}
			asciiartweb.RenderTemplate(w, "./templates/errorPage.html", a, http.StatusNotFound)
			return
		}
		http.StripPrefix("/assets/", fs).ServeHTTP(w, r)
	})

	http.HandleFunc("/", asciiartweb.IndexPage)
	http.HandleFunc("/ascii-art", asciiartweb.AsciiArtPage)
	fmt.Println("http://localhost:8080/")
	log.Fatal(http.ListenAndServe(port, nil))
}
