package assets

import (
	"embed"
	"net/http"
	"strings"
)

//go:embed *
var content embed.FS

func ReadFile(name string) ([]byte, error) {
	return content.ReadFile(name)
}

func EnsureReadFile(name string) []byte {
	data, err := ReadFile(name)
	if err != nil {
		panic(err)
	}
	return data
}

func Server() http.Handler {
	fs := http.FS(content)
	srv := http.FileServer(fs)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Cache-Control", "public, max-age=0")
		srv.ServeHTTP(w, r)
	})
}
