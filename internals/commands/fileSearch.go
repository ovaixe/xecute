package commands

import (
	"flag"
	"fmt"
	"os"

	"github.com/ovaixe/xecute/internals/search"
	"github.com/ovaixe/xecute/internals/utils"
)

type SearchCommand struct {
	CMD         *flag.FlagSet
	root        *string
	insensitive *bool
	fileName    string
	result      []string
}

func NewSearchCommand() SearchCommand {
	searchCmd := flag.NewFlagSet("s", flag.ExitOnError)
	rootDir := searchCmd.String("r", "/", "Root directory")
	insensitive := searchCmd.Bool("i", false, "Ignore case sensitive")

	searchCmd.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s: searches for filename\n", searchCmd.Name())
		fmt.Fprintf(os.Stderr, "  %s [options] <filename>\n", searchCmd.Name())
		fmt.Fprintf(os.Stderr, "Options:\n")
		searchCmd.PrintDefaults()
	}

	return SearchCommand{
		CMD:         searchCmd,
		root:        rootDir,
		insensitive: insensitive,
	}
}

func (command *SearchCommand) Execute() {
	command.CMD.Parse(os.Args[2:])
	if command.CMD.NArg() < 1 {
		fmt.Println("expected filename")
		command.CMD.Usage()
		os.Exit(1)
	}

	command.fileName = command.CMD.Arg(0)

	filePaths, err := search.SearchFile(*command.root, command.fileName, *command.insensitive)
	if err != nil {
		utils.WriteError("ERROR", "", err)
		os.Exit(1)
	}

	command.result = filePaths
	command.printResults()
}

func (command *SearchCommand) printResults() {
	fmt.Println("Results:")
	for _, res := range command.result {
		fmt.Print("File Path: ")
		fmt.Println(res)
	}
}
