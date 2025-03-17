package jumpserver

import "net/http"

type SearchPage struct {
	Query   string
	Field   string
	Results SearchResults
}

func (j *JumpServer) handleSearch(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query().Get("q")
	field := r.URL.Query().Get("field")

	var results SearchResults
	if query != "" {
		// Handle field-specific searches
		switch field {
		case "hostname":
			results = j.searchByField(query, field, func(h *Hostname) bool {
				return h.HostSearch(query)
			})
		case "address":
			results = j.searchByField(query, field, func(h *Hostname) bool {
				return h.AddrSearch(query)
			})
		case "registrant":
			results = j.searchByField(query, field, func(h *Hostname) bool {
				return h.Registrant.RegistrarSearch(query)
			})
		case "description":
			results = j.searchByField(query, field, func(h *Hostname) bool {
				return h.Registrant.TextSearch(query)
			})
		case "tags":
			results = j.searchByField(query, field, func(h *Hostname) bool {
				return h.Registrant.HasTag(query)
			})
		default:
			// Default to full search across all fields
			results = j.Search(query)
		}
	}

	page := &SearchPage{
		Query:   query,
		Field:   field,
		Results: results,
	}

	err := templates.ExecuteTemplate(w, "search.html", page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Helper function to search by specific field
func (j *JumpServer) searchByField(query string, field string, matcher func(*Hostname) bool) SearchResults {
	var results SearchResults
	for _, host := range j.Hostnames {
		if matcher(host) {
			results = append(results, &SearchResult{
				Hostname:  host,
				Host:      field == "hostname",
				Addr:      field == "address",
				Registrar: field == "registrant",
				Text:      field == "description",
				Tag:       field == "tags",
			})
		}
	}
	return results
}
