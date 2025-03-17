package jumpserver

import (
	"fmt"
	"strings"
)

// Registrant is a struct that contains the type, name and tags of the registrant.
// This is used to store the registrant metadata of the hostname.
type Registrant struct {
	Type        string   `json:"type"`        // The type of the registrant, e.g. "user" or "service"
	Name        string   `json:"name"`        // The name of the registrant, either the username of the registrant or the hostname of the service
	Tags        []string `json:"tags"`        // The tags of the registry, e.g. "admin", "web", "game" etc.
	Description string   `json:"description"` // The description of the registrant obtained from the meta-tags provided by the page, if any
}

func (r *Registrant) String() string {
	return fmt.Sprintf("%s:%s:%s:%s", r.Type, r.Name, r.Tags, r.Description)
}

func (r *Registrant) ByUser() bool {
	if r.Type == "user" {
		return true
	}
	return false
}

func (r *Registrant) ByService() bool {
	if r.Type == "service" {
		return true
	}
	return false
}

func (r *Registrant) HasTag(tag string) bool {
	for _, t := range r.Tags {
		if t == tag {
			return true
		}
	}
	return false
}

func (r *Registrant) DataSearch(query string) (bool, bool, bool) {
	registrar := r.RegistrarSearch(query)
	text := r.TextSearch(query)
	tag := r.HasTag(query)
	return registrar, text, tag
}

func (r *Registrant) TextSearch(query string) bool {
	if strings.Contains(r.Description, query) {
		return true
	}
	return false
}

func (r *Registrant) RegistrarSearch(query string) bool {
	if strings.Contains(r.Name, query) {
		return true
	}
	return false
}
