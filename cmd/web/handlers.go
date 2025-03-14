package main

import (
	"errors"
	"fmt"
	"github.com/pixfloage/snippetbox/internal/models"
	"github.com/pixfloage/snippetbox/internal/validator"
	"net/http"
	"strconv"
)

type snippetCreateForm struct {
	Title                   string `form:"title"` // map from HTML form "title" to this struct Title
	Content                 string `form:"content"`
	Expires                 int    `form:"expires"` // struct embedding, ignore this field during decoding
	validator.FormValidator `form:"-"`
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// w.Header().Add("Server", "Go") // set by a middleware instead

	snippets, err := app.snippets.Latest()

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.newTemplateData(r)
	data.Snippets = snippets

	app.render(w, r, http.StatusOK, "home.html", data)
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}

		return
	}

	data := app.newTemplateData(r)
	data.Snippet = snippet

	app.render(w, r, http.StatusOK, "view.html", data)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	data.Form = snippetCreateForm{
		Expires: 365,
	}

	app.render(w, r, http.StatusOK, "create.html", data)
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	var form snippetCreateForm
	if err := app.decodePostForm(r, &form); err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.CheckField(validator.NotBlank(form.Title), "title", "Title cannot be blank")
	form.CheckField(validator.MaxChars(form.Title, 180), "title", "Title cannot be more than 180 characters")

	form.CheckField(validator.NotBlank(form.Content), "content", "Content cannot be blank")

	form.CheckField(validator.PermittedValue(form.Expires, 1, 7, 365), "expires", "Expires could only be 1, 7 or 365 days")

	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, r, http.StatusUnprocessableEntity, "create.html", data)
		return
	}

	id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
