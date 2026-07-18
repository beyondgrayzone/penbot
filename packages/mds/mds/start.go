package mds

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var showVersion = flag.Bool("version", false, "Show version")
var inFilePath = flag.String("in", "", "File path to raw markdown file")
var outFilePath = flag.String("out", "", "File path to processed output markdown file")

func Main() {
	// Check for subcommands before flag.Parse() so flag doesn't consume them.
	if len(os.Args) > 1 {
		subcmd := os.Args[1]
		// If it starts with "-" it's a flag for the parent, otherwise it's a subcommand.
		if !strings.HasPrefix(subcmd, "-") {
			switch subcmd {
			case "version":
				HandleVersionSubcommand(os.Args[2:])
				return
			case "sections":
				HandleSectionsSubcommand(os.Args[2:])
				return
			case "theme":
				HandleThemeSubcommand(os.Args[2:])
				return
			case "help":
				Usage()
				return
			default:
				fmt.Fprintf(os.Stderr, "Error: unknown command %q.\n\n", subcmd)
				Usage()
				os.Exit(1)
			}
		}
	}

	flag.CommandLine.SetOutput(os.Stdout)
	flag.Usage = Usage
	flag.Parse()

	if len(*inFilePath) == 0 || len(*outFilePath) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	err := Start(*inFilePath, *outFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Println("Markdown processing a success")
}
