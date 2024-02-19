package htmlutil

import (
	"html/template"
	log "log/slog"
	"net/http"
)

var (
	tpl *template.Template
)

func init() {
	tpl, _ = template.ParseGlob("frontend/templates/*.html")
}

func RenderTemplate(w http.ResponseWriter, templateName string, data interface{}) {
	err := tpl.ExecuteTemplate(w, templateName, data)
	if err != nil {
		http.Error(w, "Failed to render template: "+err.Error(), http.StatusInternalServerError)
		log.Error("Error rendering template: %s", err)
	}
}
