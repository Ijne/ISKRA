package tools

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) error {
	templates := template.Must(template.ParseFiles(
		//"internal/templates/base.html",
		"internal/templates/" + tmpl,
	))

	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		return fmt.Errorf("error with execute template: %s", err)
	}
	return nil
}
