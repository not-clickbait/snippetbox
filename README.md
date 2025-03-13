# Snippetbox
A pet project for learning [Go](https://go.dev/) based off [Let's Go - Alex Edwards](https://lets-go.alexedwards.net/) book.

## Features
TBD

## Setup
### Locally
For now, the DSN is hardcoded to be `web:web@/snippetbox?parseTime=true`, a MySQL with user `web` and password `web` is expected.

To spin up the http server, run `go run ./cmd/web` from the root directory. 

Open `http://localhost:4000` in a browser to get started with Snippetbox.