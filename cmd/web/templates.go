package main

import (
	"bytes"
	"fmt"
	"github.com/pixfloage/snippetbox/internal/models"
	"html/template"
	"net/http"
	"path/filepath"
	"time"
)

// Create a humanDate function which returns a nicely formatted string // representation of a time.Time object.
func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

type templateData struct {
	CurrentYear int
	Snippet     models.Snippet
	Snippets    []models.Snippet
	Form        any
	Flash       string
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.html")

	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		// The template.FuncMap must be registered with the template set before you
		// call the ParseFiles() method. This means we have to use template.New() to
		// create an empty template set, use the Funcs() method to register the
		// template.FuncMap, and then parse the file as normal.
		ts := template.New(name).Funcs(functions)

		// Base
		ts, err = ts.ParseFiles("./ui/html/base.html")
		if err != nil {
			return nil, err
		}

		// Add partials to the Base template set
		ts, err = ts.ParseGlob("./ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}

		//files := []string{
		//	"../../ui/html/base.html",
		//	"../../ui/html/partials/nav.html",
		//	page,
		//}

		// The actual page template
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data templateData) {
	ts, ok := app.templateCache[page]

	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, r, err)
		return
	}

	// Write to a buffer first to check rendering runtime errors, only respond with html if it rendered without error
	buf := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	w.WriteHeader(status)

	buf.WriteTo(w)

	//if err := ts.ExecuteTemplate(w, "base", data); err != nil {
	//	app.serverError(w, r, err)
	//	return
	//}
}
