package jumpserver

import (
	gohtmlmetadata "github.com/go-i2p/go-html-metadata"
	"github.com/go-i2p/onramp"
)

type JumpServer struct {
	*gohtmlmetadata.Extractor
	Index     string      `json:"index"`     // The intro page/index page content of the jump server
	Hostnames []*Hostname `json:"hostnames"` // The hostnames of the jump server
	SyncURLs  []string    `json:"syncurls"`  // The URLs to sync the hostnames of the jump server with
	Garlic    *onramp.Garlic
}

func (j *JumpServer) AddHostname(h *Hostname) {
	j.Hostnames = append(j.Hostnames, h)
}

func (j *JumpServer) RemoveHostname(h *Hostname) {
	for i, host := range j.Hostnames {
		if host == h {
			j.Hostnames = append(j.Hostnames[:i], j.Hostnames[i+1:]...)
		}
	}
}

type SearchResult struct {
	*Hostname
	Registrar bool
	Text      bool
	Tag       bool
	Addr      bool
	Host      bool
}

type SearchResults []*SearchResult

func (j *JumpServer) Search(query string) SearchResults {
	var results SearchResults
	for _, host := range j.Hostnames {
		registrar, text, tag, addr, hostname := host.FullSearch(query)
		if registrar || text || tag || addr || hostname {
			results = append(results, &SearchResult{
				Hostname:  host,
				Registrar: registrar,
				Text:      text,
				Tag:       tag,
				Addr:      addr,
				Host:      hostname,
			})
		}
	}
	return results
}
