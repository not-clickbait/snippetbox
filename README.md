# Snippetbox
A pet project for learning [Go](https://go.dev/) based off [Let's Go - Alex Edwards](https://lets-go.alexedwards.net/) book.

## Features
TBD

## Setup
### Locally
For now, the DSN is hardcoded to be `web:web@/snippetbox?parseTime=true`, a MySQL with user `web` and password `web` is expected.

To spin up the http server, run `go run ./cmd/web` from the root directory. 

Open `http://localhost:4000` in a browser to get started with Snippetbox.

### Self-signed TLS Generation
Locate GO standard library location 
`/usr/local/go/src/crypto/tls` (non Homebrew)

Run `generate_cert.go` inside the /tls folder
```shell
go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```