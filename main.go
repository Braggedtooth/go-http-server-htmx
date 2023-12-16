package main

import (
	"braggedtooth/pkg"
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.ServeFile(w, r, "../templates/hello.html")
			return
		}

		name := r.FormValue("name")
		message, err := pkg.Greet(name)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Message string
		}{
			Message: message,
		}

		tmpl, err := template.ParseFiles("../templates/hello.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
