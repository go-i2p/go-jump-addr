package jumpserver

import (
	"crypto/tls"
	"net/http"

	"github.com/go-i2p/onramp"

	gohtmlmetadata "github.com/go-i2p/go-html-metadata"
)

func NewServer() (*JumpServer, error) {
	garlic, err := onramp.NewGarlic("jumpserver", "127.0.0.1:7656", onramp.OPT_DEFAULTS)
	if err != nil {
		return nil, err
	}
	roundTripper := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		DialContext: garlic.DialContext,
	}
	extractor := gohtmlmetadata.NewExtractor(roundTripper)
	return &JumpServer{
		Extractor: extractor,
		Hostnames: make([]*Hostname, 0),
		Garlic:    garlic,
	}, nil
}
