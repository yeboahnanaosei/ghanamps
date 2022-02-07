// This is the cli tool for ghanamps. It returns a JSON array containing the
// details of all the current members of Ghana's parliament.
package main

import (
	"flag"
	"fmt"
	"os"
)

const command = "ghanamps"

var (
	// ldflags can be ued to set the version string.
	version       = "development"
	arch          = "unknown architecture"

	versionString = fmt.Sprintf("%s version %s %s", command, version, arch)
	userAgent     = fmt.Sprintf("%s/%s (%s) A cli tool for getting information about Ghana's parliament", command, version, arch)
)

func main() {
	members := flag.NewFlagSet("members", flag.ExitOnError)
	party := members.String("party", "", "filter members by party")
	leadership := flag.NewFlagSet("leadership", flag.ExitOnError)
	flag.Usage = usage
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "%s: missing subcommand: See `%s -h` for help\n", command, os.Args[0])
		os.Exit(1)
	}

	switch os.Args[1] {
	case "members":
		members.Parse(os.Args[2:])
		handleGetMembers(*party)
	case "leaders":
		leadership.Parse(os.Args[2:])
		handleGetLeadership()
	case "version":
		fmt.Fprint(os.Stdout, versionString)
		os.Exit(0)
	default:
		fmt.Fprintf(os.Stderr, "%s: unknown subcommand '%s'. See %s -h\n", command, os.Args[1], os.Args[0])
		os.Exit(1)
	}
}
