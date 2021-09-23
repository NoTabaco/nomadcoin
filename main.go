package main

import (
	"github.com/NoTabaco/nomadcoin/explorer"
	"github.com/NoTabaco/nomadcoin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
