// This is the cli tool for ghanamps. It returns a JSON array containing the
// details of all the current members of Ghana's parliament.
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/yeboahnanaosei/ghanamps"
)

func main() {
	members, err := ghanamps.Fetch()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ghanamps: error fetching members: %s", err)
		os.Exit(1)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetEscapeHTML(false)
	enc.SetIndent(" ", "  ")
	if err := enc.Encode(members); err != nil {
		fmt.Fprintf(os.Stderr, "ghanamps: failed to fetch members: %s", err)
	}
}
