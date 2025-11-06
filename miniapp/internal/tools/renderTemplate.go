package tools

import (
	"fmt"
	"html/template"
	"iskra/shared/config"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, cfg *config.Config, tmpl string, data interface{}) error {
	templates := template.Must(template.ParseFiles(
		//"internal/templates/base.html",
		cfg.MiniApp.TemplatesPath + tmpl,
	))

	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		return fmt.Errorf("error with execute template: %s", err)
	}
	return nil
}
