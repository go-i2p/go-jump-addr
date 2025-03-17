package jumpserver

import (
	"log"
	"net/http"
	"strings"
	"time"
)

func (j *JumpServer) Serve() error {
	l, err := j.Garlic.Listen()
	if err != nil {
		return err
	}
	log.Printf("Listening on %s\n", l.Addr())
	defer l.Close()
	return http.Serve(l, j)
}

func (j *JumpServer) Close() error {
	return j.Garlic.Close()
}

func (j *JumpServer) updateMetadata() error {
	// Update the metadata of the jump server
	for _, host := range j.Hostnames {
		metadata, err := j.Extractor.Extract(host.I2PAddr.Base32())
		if err != nil {
			return err
		}
		for _, tag := range metadata {
			if host.Registrant.Description != "" {
				continue
			}
			if host.Time == (time.Time{}) {
				host.Time = time.Now()
			}
			if tag.Name == "title" {
				host.Registrant.Description = tag.Content
			}
			if tag.Name == "author" {
				host.Registrant.Description = tag.Content
			}
			if tag.Name == "description" {
				host.Registrant.Description = tag.Content
			}
			if tag.Name == "keywords" {
				for _, keyword := range splitTags(tag.Content) {
					host.Registrant.Tags = append(host.Registrant.Tags, keyword)
				}
			}

		}
	}
	return nil
}

func splitTags(tags string) []string {
	// replace commas with spaces
	tags = strings.Replace(tags, ",", " ", -1)
	// split the tags
	split := strings.Split(tags, " ")
	return split
}
