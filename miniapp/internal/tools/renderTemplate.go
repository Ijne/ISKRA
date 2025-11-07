package tools

import (
	"fmt"
	"html/template"
	"iskra/shared/config"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, cfg *config.Config, tmpl string, data interface{}) error {
	var templates *template.Template
	if tmpl != "start.html" {
		templates = template.Must(template.ParseFiles(
			cfg.MiniApp.TemplatesPath+"base.html",
			cfg.MiniApp.TemplatesPath+tmpl,
		))
	} else {
		templates = template.Must(template.ParseFiles(
			cfg.MiniApp.TemplatesPath + tmpl,
		))
	}

	if tmpl != "start.html" {
		err := templates.ExecuteTemplate(w, "base.html", data)
		if err != nil {
			return fmt.Errorf("error with execute template: %s", err)
		}
	} else {
		err := templates.ExecuteTemplate(w, "start.html", data)
		if err != nil {
			return fmt.Errorf("error with execute template: %s", err)
		}
	}
	return nil
}
