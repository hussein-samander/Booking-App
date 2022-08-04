package render

import (
	"bytes"
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/hussein-samander/Booking-App/config"
	"github.com/hussein-samander/Booking-App/models"
)

var functions = make(template.FuncMap)

var app *config.AppConfig

func NewConfig(a *config.AppConfig) {
	app = a
}

func AddData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplates(w http.ResponseWriter, temp string, td *models.TemplateData) {
	tc := make(map[string]*template.Template)
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	parsedTemplate, ok := tc[temp]
	if !ok {
		fmt.Fprintf(w, "No such page")
		return
	}
	//Why not directly parsedTemplate.Execute(w, nil)?
	//Guy says it is not being read from desk this time
	buf := new(bytes.Buffer)
	parsedTemplate.Execute(buf, td)
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

// CreateTemplateCache creates a template cache as a template
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := make(map[string]*template.Template)
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}
		if len(matches) != 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
