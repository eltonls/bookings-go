package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/eltonls/bookings/pkg/config"
	"github.com/eltonls/bookings/pkg/models"
)

var app *config.AppConfig

func SetAppConfig(appConfig *config.AppConfig) {
	app = appConfig
}

func SetDefaultData(templateData *models.TemplateData) *models.TemplateData {
	
	return templateData
}

func RenderTemplate(w http.ResponseWriter, tmpl string, templateData *models.TemplateData) {
	var templateCache map[string]*template.Template

	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	template, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("Could not get the template")
	}

	buf := new(bytes.Buffer)

	templateData = SetDefaultData(templateData)

	err := template.Execute(buf, templateData)
	if err != nil {
		log.Println(err)
	}

	// Render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}



func CreateTemplateCache() (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return templateCache, err
	}

	for _, page := range pages {
		filename := filepath.Base(page)
		templateSet, err := template.New(filename).ParseFiles(page)
		if err != nil {
			return templateCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return templateCache, err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return templateCache, err
			}
		}

		templateCache[filename] = templateSet
	}

	return templateCache, nil
}



















/* VERSION FOR LITTLE BABYES */

// var templateCache = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, templateName string) {
// 	var tmpl *template.Template
// 	var err error
//
// 	// Check to see if the template is already cached
// 	_, inMap := templateCache[templateName]
// 	if inMap {
// 		// We have the template in cache
// 		log.Println("Using cached template")
// 	} else {
// 		// we Don't have the template in cache, need to create it
// 		fmt.Println("Creating template and adding it to cache")
// 		err = createTemplateCache(templateName)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	}
//
// 	tmpl = templateCache[templateName]
//
// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }
//
// func createTemplateCache(templateName string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", templateName),
// 		"./templates/base.layout.tmpl",
// 	}
//
// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}
//
// 	templateCache[templateName] = tmpl
//
// 	return nil
// }
