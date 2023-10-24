package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./templates/" + tmpl)
	_ = parsedTemplate.Execute(w, nil)
	// _ := parsedTemplate.Execute(w, nil)
	fmt.Println("Error rendering template:", err)
	if err != nil {
		fmt.Println("Error parsing template:", err)
	}
}
