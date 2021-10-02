package main

import (
	"github.com/NoTabaco/nomadcoin/blockchain"
	"github.com/NoTabaco/nomadcoin/cli"
)

func main() {
	blockchain.Blockchain()
	cli.Start()
}
