module github.com/go-i2p/go-jump-addr

go 1.26.0

require (
	github.com/go-i2p/go-html-metadata v0.0.0-20260908192746-ec07a7866a6f
	github.com/go-i2p/i2pkeys v0.33.92
	github.com/go-i2p/onramp v0.33.92
)

require (
	github.com/cretz/bine v0.2.0 // indirect
	github.com/go-i2p/sam3 v0.33.92 // indirect
	github.com/sirupsen/logrus v1.10.2 // indirect
	golang.org/x/crypto v0.57.0 // indirect
	golang.org/x/net v0.58.0 // indirect
	golang.org/x/sys v0.48.0 // indirect
)

retract (
	v0.1.59999
	v0.1.5999
	v0.1.599
)
