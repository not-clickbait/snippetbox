package main

import (
	"github.com/justinas/alice"
	"github.com/pixfloage/snippetbox/internal/nfs"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	dynamic := alice.New(app.sessionManager.LoadAndSave)

	// fileServer := http.FileServer(http.Dir("./ui/static/"))
	fileServer := http.FileServer(&nfs.NeuteredFileSystem{Fs: http.Dir("./ui/static/")})
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.Handle("GET /{$}", dynamic.ThenFunc(app.home)) // terminate / with ${$} so it won't act as a catch-all
	mux.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(app.snippetView))
	mux.Handle("GET /snippet/create", dynamic.ThenFunc(app.snippetCreate))
	mux.Handle("POST /snippet/create", dynamic.ThenFunc(app.snippetCreatePost))

	// return mux

	// use the commonHeaders middleware
	// return app.recoverPanic(app.logRequest(commonHeaders(mux)))
	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)
	return standard.Then(mux)
}
