package main

import (
	"github.com/NoTabaco/nomadcoin/cli"
	"github.com/NoTabaco/nomadcoin/db"
)

func main() {
	cli.Start()
	defer db.Close()
}
