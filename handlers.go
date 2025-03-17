package jumpserver

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var templates *template.Template

func init() {
	templates = template.Must(template.ParseFiles(
		filepath.Join("tpl", "index.html"),
		filepath.Join("tpl", "add.html"),
	))
}

func (j *JumpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/", "/index.html":
		err := templates.ExecuteTemplate(w, "index.html", j)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case "/add", "/add.html":
		j.handleAdd(w, r)
	case "/search":
		j.handleSearch(w, r)
	case "/all-hosts.txt":
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(j.HostsFile()))
	case "/static/style.css":
		http.ServeFile(w, r, filepath.Join("static", "style.css"))
	case "/static/script.js":
		http.ServeFile(w, r, filepath.Join("static", "script.js"))
	default:
		http.NotFound(w, r)
	}
}
