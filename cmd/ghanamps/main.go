// This is the cli tool for ghanamps. It returns a JSON array containing the
// details of all the current members of Ghana's parliament.
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	members := flag.NewFlagSet("members", flag.ExitOnError)
	party := members.String("party", "", "filter members by party")
	leadership := flag.NewFlagSet("leadership", flag.ExitOnError)
	flag.Usage = usage
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "ghanamps: missing subcommand: See `%s -h` for help\n", os.Args[0])
		os.Exit(1)
	}

	switch os.Args[1] {
	case "members":
		members.Parse(os.Args[2:])
		handleGetMembers(*party)
	case "leaders":
		leadership.Parse(os.Args[2:])
		handleGetLeadership()
	default:
		fmt.Fprintf(os.Stderr, "ghanamps: unknown subcommand '%s'. See %s -h\n", os.Args[1], os.Args[0])
		os.Exit(1)
	}
}
