package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
)

// addStaticFileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem. It takes a router and then two other
// parameters: the path prefix for the URLs at which static assets should be
// served and the directory from which we should serve them. So, you'll be
// able to reference assets like "/static/foo.css" and it will serve from
// "staticfiles/foo.css" if both `path` is "static" `staticFileDir` is
// "staticfiles".
func addStaticFileServer(r chi.Router, path string, staticFileDir string) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}
	workDir, _ := os.Getwd()
	root := http.Dir(filepath.Join(workDir, staticFileDir))

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(path)
		fs.ServeHTTP(w, r)
	}))
}
