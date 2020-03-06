package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func parseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

func index(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("templates/index.html"))
	t.Execute(writer, "")
}

func login(writer http.ResponseWriter, request *http.Request) {
	t := parseTemplateFiles("login.layout", "public.navbar", "login")
	t.Execute(writer, nil)
}
