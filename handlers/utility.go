package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Printf("Error while parsing template: %v", err)
	}
	err = t.Execute(w, templateData)
	if err != nil {
		log.Printf("Error while execute template: %v", err)
	}
}
