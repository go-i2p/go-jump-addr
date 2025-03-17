package jumpserver

import (
	"html/template"
	"net/http"
	"path/filepath"
	"time"

	"github.com/go-i2p/i2pkeys"
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

func (j *JumpServer) handleAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := templates.ExecuteTemplate(w, "add.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		// Extract form data
		hostname := r.Form.Get("hostname")
		destination := r.Form.Get("destination")
		regType := r.Form.Get("type")
		name := r.Form.Get("name")
		description := r.Form.Get("description")
		tags := r.Form.Get("tags")

		// Validate hostname
		if hostname == "" {
			http.Error(w, "Hostname is required", http.StatusBadRequest)
			return
		}

		// Parse I2P address
		addr, err := i2pkeys.NewI2PAddrFromString(destination)
		if err != nil {
			http.Error(w, "Invalid I2P destination", http.StatusBadRequest)
			return
		}

		// Create new hostname entry
		host := &Hostname{
			I2PAddr: &addr,
			Time:    time.Now(),
			Registrant: Registrant{
				Type:        regType,
				Name:        name,
				Description: description,
				Tags:        splitTags(tags),
			},
			Hostname: hostname,
		}

		// Add to jump server
		j.AddHostname(host)

		// Redirect to index page
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func (j *JumpServer) HostsFile() string {
	var hosts string
	for _, host := range j.Hostnames {
		hosts += host.Hostname + "=" + host.Base64() + "\n"
	}
	return hosts
}
