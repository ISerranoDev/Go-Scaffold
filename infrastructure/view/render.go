package view

import (
	"net/http"
)

var TemplateCacheMap TemplateCache

func Render(w http.ResponseWriter, name string, data interface{}) {
	t, ok := TemplateCacheMap[name]
	if !ok {
		http.Error(w, "Plantilla no encontrado: "+name, http.StatusInternalServerError)
		return
	}

	err := t.ExecuteTemplate(w, name, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
