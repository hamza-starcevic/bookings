package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/hamza-starcevic/bookings/pkg/config"
	"github.com/hamza-starcevic/bookings/pkg/models"
)

var app *config.AppConfig

// * NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}
func AddDefaultData(td *models.Templatedata) *models.Templatedata {
	return td
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.Templatedata) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	//get the template cache from app config

	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	//get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	buff := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buff, td)
	if err != nil {
		log.Println(err)
	}

	//render the template
	_, err = buff.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	//get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through all files ending in *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}

	return myCache, nil
}
