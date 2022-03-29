package web

import (
	"fmt"
	"html/template"
	"net/http"

	"go-short/shortener"
	"go-short/storage"
)

var tmpl = template.Must(template.ParseGlob("views/*.html"))
var urlShortener = shortener.New(storage.New("localhost:6379"))

func renderPage(w http.ResponseWriter, template string, content interface{}) {
	if err := tmpl.ExecuteTemplate(w, template, content); err != nil {
		http.Error(w, fmt.Sprintf("rendering template: %v", err), http.StatusInternalServerError)
		return
	}
}

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		if r.Method == http.MethodPost {
			createShortenedURL(w, r)
		} else if r.Method == http.MethodGet {
			renderPage(w, "home.html", nil)
		}
	} else if r.Method == http.MethodGet {
		handleRedirect(w, r)
	}
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	url, err := urlShortener.Get(key)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to find %s: %v", key, err), http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func createShortenedURL(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if r.Form["url"] == nil || len(r.Form["url"]) == 0 {
		http.Error(w, "no url provided", http.StatusBadRequest)
		return
	}

	url := r.Form["url"][0]
	shortID, err := urlShortener.Shorten(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to shorten %s: %v", url, err), http.StatusInternalServerError)
		return
	}

	renderPage(w, "created.html", struct {
		ShortID string
	}{shortID})
}

func RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", handleHomePage)
}
