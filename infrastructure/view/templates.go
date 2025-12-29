package view

import (
	"html/template"
	"log"
	"path/filepath"
)

type TemplateCache map[string]*template.Template

func NewTemplateCache() (TemplateCache, error) {
	cache := TemplateCache{}

	pages, err := filepath.Glob("views/*.gohtml")
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		// layouts opcionales
		layouts, err := filepath.Glob("views/layouts/*.gohtml")
		if err == nil && len(layouts) > 0 {
			ts, err = ts.ParseGlob("views/layouts/*.gohtml")
			if err != nil {
				return cache, err
			}
		}

		cache[name] = ts
		log.Println("Template cargado:", name)
	}

	return cache, nil
}
