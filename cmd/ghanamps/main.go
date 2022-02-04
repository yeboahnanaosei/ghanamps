// This is the cli tool for ghanamps. It returns a JSON array containing the
// details of all the current members of Ghana's parliament.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/yeboahnanaosei/ghanamps"
)

func main() {
	fetchMembers := flag.Bool("members", false, "returns the current members of Ghana's parliament")
	fetchLeaders := flag.Bool("leaders", false, "returns the current leaders of Ghana's parliament")
	flag.Parse()

	if !*fetchMembers && !*fetchLeaders {
		fmt.Fprint(os.Stderr, "ghanamps: you will have to specify one of -members or -leaders\n")
		flag.Usage()
		os.Exit(1)
	}
	if *fetchMembers && *fetchLeaders {
		fmt.Fprint(os.Stdout, "ghanapms: please specify only one of the flags: members or leaders\n")
		flag.Usage()
		os.Exit(1)
	}

	var payload interface{}
	var err error

	if *fetchMembers {
		payload, err = ghanamps.Members()
		if err != nil {
			fmt.Fprintf(os.Stderr, "ghanamps: error fetching members: %s\n", err)
			os.Exit(1)
		}
	}

	if *fetchLeaders {
		payload, err = ghanamps.Leaders()
		if err != nil {
			fmt.Fprintf(os.Stderr, "ghanamps: error fetching leaders: %s\n", err)
			os.Exit(1)
		}
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetEscapeHTML(false)
	enc.SetIndent(" ", "  ")
	if err := enc.Encode(payload); err != nil {
		fmt.Fprintf(os.Stderr, "ghanamps: failed to fetch leaders: %s", err)
	}
}
