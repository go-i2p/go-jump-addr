package jumpserver

import (
	"context"
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	j.AutoUpdateMetadata(ctx)
	j.StartSync(j.SyncURLs, ctx)
	return http.Serve(l, j)
}

func (j *JumpServer) Close() error {
	return j.Garlic.Close()
}

func (j *JumpServer) AutoUpdateMetadata(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				if err := j.updateMetadata(); err != nil {
					log.Printf("Error updating metadata: %v", err)
				}
			}
		}
	}()
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
