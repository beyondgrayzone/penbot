package mds

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func isHelpArg(arg string) bool {
	if !strings.HasPrefix(arg, "-") {
		return false
	}
	arg = strings.TrimPrefix(arg[1:], "-")
	return arg == "h" || arg == "help"
}
func hasHelpFlag(args []string) bool {
	for _, arg := range args {
		if isHelpArg(arg) {
			return true
		}
	}
	return false
}

func usage(s string) {
	f := flag.CommandLine.Output()
	fmt.Fprintf(f, "%s\n", s)
	if hasHelpFlag(os.Args[1:]) {
		flag.PrintDefaults()
	} else {
		fmt.Fprintf(f, `Run "%s -help" in order to see the description for all the available flags`+"\n", os.Args[0])
	}
}

func Usage() {
	const s = `
Help:
  mds -in <input-file> -out <output-file>
      Process a raw markdown file and write sanitized output.

  mds version add <version> [--copy]
      Add a new documentation version (e.g. v3) with route scaffolding.
      Use --copy to copy mds content from the previous version.

  mds version remove <version>
      Remove a documentation version and its route scaffolding.

  mds version sync
      Sync all versions: ensure routes, mds, and search.json are in sync.

  mds sections
      List, add, or remove documentation sections in docs/velite.config.js.
      Run "mds sections help" for subcommand usage.

  mds theme
      List available themes or apply a new default theme.
      Run "mds theme help" for subcommand usage.

  mds help
      Show this help message.
`
	usage(s)
}
