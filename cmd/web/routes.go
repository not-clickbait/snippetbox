package main

import (
	"github.com/justinas/alice"
	"github.com/pixfloage/snippetbox/internal/nfs"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	// fileServer := http.FileServer(http.Dir("./ui/static/"))
	fileServer := http.FileServer(&nfs.NeuteredFileSystem{Fs: http.Dir("./ui/static/")})
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home) // terminate / with ${$} so it won't act as a catch-all
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	// return mux

	// use the commonHeaders middleware
	// return app.recoverPanic(app.logRequest(commonHeaders(mux)))
	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)
	return standard.Then(mux)
}
