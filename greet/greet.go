package greet

import (
	"errors"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
)

func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
/* func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}

	return messages, nil
} */

func Hellos(names []string) ([]map[string]string, error) {
	messages := make([]map[string]string, len(names))

	for i, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[i] = map[string]string{name: message}
	}

	return messages, nil
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
