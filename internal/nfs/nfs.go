package nfs

import (
	"net/http"
	"path/filepath"
)

/*
Disables directory listings for the http.FileServer
https://www.alexedwards.net/blog/disable-http-fileserver-directory-listings
*/

type NeuteredFileSystem struct {
	Fs http.FileSystem
}

func (nfs *NeuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.Fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()

	if err != nil {
		return nil, err
	}

	if s.IsDir() {
		index := filepath.Join(path, "index.html")

		// This calls the FS Open func with the index.html path, if it exists, it returns the file
		if _, err = nfs.Fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}
			return nil, err
		}
	}

	return f, nil
}
