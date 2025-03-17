package main

import jumpserver "github.com/go-i2p/go-jump-addr"

func main() {
	if js, err := jumpserver.NewServer(); err != nil {
		panic(err)
	} else {
		if err := js.Serve(); err != nil {
			panic(err)
		}
	}
}
