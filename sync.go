package jumpserver

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/go-i2p/i2pkeys"
)

func (j *JumpServer) StartSync(urls []string, ctx context.Context) {
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				for _, url := range urls {
					j.syncHostnames(url)
				}
			}
		}
	}()
}

func (j *JumpServer) syncHostnames(u string) (host, content string) {
	uri, err := url.Parse(u)
	if err != nil {
		log.Printf("Failed to parse URL: %s\n", err)
		return "", ""
	}
	if uri.Scheme == "" {
		uri.Scheme = "http"
	}
	if uri.Host == "" {
		log.Printf("Failed to parse URL: no host\n")
		return "", ""
	}
	// Sync the hostnames of the jump server
	// with the hostnames of the jump server
	// at the given URL
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		log.Printf("Failed to create request: %s\n", err)
		return "", ""
	}
	req.Header.Set("User-Agent", "go-jump-addr")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Failed to get response: %s\n", err)
		return "", ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Printf("Failed to get response: %s\n", resp.Status)
		return "", ""
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response: %s\n", err)
		return "", ""
	}
	// Extract the hostnames from the response
	// and add them to the jump server
	for _, host := range extractHostnames(body) {
		vals := strings.Split(host, "=")
		if len(vals) != 2 {
			log.Printf("Failed to parse hostname: %s\n", host)
			continue
		}
		hostname, dest := vals[0], vals[1]
		addr, err := i2pkeys.NewI2PAddrFromString(dest)
		if err != nil {
			log.Printf("Failed to parse I2P address: %s\n", err)
			continue
		}
		entry := &Hostname{
			I2PAddr: &addr,
			Time:    time.Now(),
			Registrant: Registrant{
				Type:        "service",
				Name:        uri.Host,
				Description: "",
				Tags:        []string{uri.Host, uri.Scheme},
			},
			Hostname: hostname,
		}
		j.AddHostname(entry)
	}
	return uri.Host, string(body)
}

func extractHostnames(body []byte) []string {
	// Extract hostnames from the response body
	// and return them as a slice of strings
	// This function is used by SyncHostnames
	// to extract hostnames from the response
	// of the jump server at the given URL
	return strings.Split(string(body), "\n")
}
