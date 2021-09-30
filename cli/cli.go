package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/NoTabaco/nomadcoin/explorer"
	"github.com/NoTabaco/nomadcoin/rest"
)

func usage() {
	fmt.Printf("Welcome to Nomadcoin\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port:		Set the PORT of the server\n")
	fmt.Printf("-mode:		Choose one of the three options ('html', 'rest' and 'all')\n\n")
	os.Exit(0)
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	hPort := flag.Int("hPort", 3000, "Set port of the HTML Explorer")
	rPort := flag.Int("rPort", 4000, "Set port of the REST API")
	mode := flag.String("mode", "rest", "Choose one of the three options ('html', 'rest' and 'all')")

	flag.Parse()

	switch *mode {
	case "rest":
		{
			rest.Start(*rPort)
		}
	case "html":
		{
			explorer.Start(*hPort)
		}
	case "all":
		{
			go explorer.Start(*hPort)
			rest.Start(*rPort)
		}
	default:
		{
			usage()
		}
	}
}
