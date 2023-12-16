package pkg

import (
	"errors"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
)

func Greet(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}

func RenderGreeting(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "hello.html")
		return
	}

	name := r.FormValue("name")
	message := fmt.Sprintf("Hello, %s!", name)

	data := struct {
		Message string
	}{
		Message: message,
	}

	tmpl, err := template.ParseFiles("hello.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
