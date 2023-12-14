package main

import (
	"html/template"
	"log"
	"net/http"

	"example.com/greet"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.ServeFile(w, r, "../templates/hello.html")
			return
		}

		name := r.FormValue("name")
		message, err := greet.Hello(name)

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
		/* 		names := []string{"Gladys", "Samantha", "Darrin"}
		   		messages, err := greet.Hello(names[len(names)-1])
		   		if err != nil {
		   			log.Fatal(err)
		   		}

		   		jsonBytes, err := json.Marshal(messages)
		   		if err != nil {
		   			return
		   		}

		   		fmt.Fprint(w, string(jsonBytes)) */
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
