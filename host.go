package jumpserver

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-i2p/i2pkeys"
)

// Hostname is a struct that contains an I2PAddr and a hostname.
// This is used to store the hostnames of the jump server.
type Hostname struct {
	*i2pkeys.I2PAddr `json:"i2paddr"`    // The I2PAddr of the hostname
	time.Time        `json:"time"`       // The time the hostname was added to the jump server
	Registrant       `json:"registrant"` // The registrant of the hostname
	Hostname         string              `json:"hostname"` // The hostname of the jump server
}

func (h *Hostname) String() string {
	return fmt.Sprintf("%s:%s:%s", h.Hostname, h.I2PAddr.Base32(), h.Registrant)
}

func (h *Hostname) FullSearch(query string) (bool, bool, bool, bool, bool) {
	registrar, text, tag := h.Registrant.DataSearch(query)
	addr := h.AddrSearch(query)
	host := h.HostSearch(query)
	return registrar, text, tag, addr, host
}

func (h *Hostname) AddrSearch(query string) bool {
	if strings.Contains(h.Base32(), query) {
		return true
	}
	return false
}

func (h *Hostname) HostSearch(query string) bool {
	if strings.Contains(h.Hostname, query) {
		return true
	}
	return false
}
