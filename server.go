package jumpserver

type JumpServer struct {
	Hostnames []*Hostname `json:"hostnames"` // The hostnames of the jump server
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

func (j *JumpServer) Search(query string) []*Hostname {
	var hosts []*Hostname
	for _, host := range j.Hostnames {
		registrar, text, tag, addr, hostname := host.FullSearch(query)
		if registrar || text || tag || addr || hostname {
			hosts = append(hosts, host)
		}
	}
	return hosts
}
