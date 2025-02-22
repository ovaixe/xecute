package main

import (
	"flag"
	"fmt"
	"os"

  "github.com/ovaixe/xecute/internals/search"
)


type SearchCommand struct {
	cmd         *flag.FlagSet
	root        *string
	insensitive *bool
	fileName    string
	result      []string
}

func NewSearchCommand() SearchCommand {
	searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
	rootDir := searchCmd.String("root", "/", "Root directory")
	insensitive := searchCmd.Bool("i", false, "Ignore case sensitive")

	searchCmd.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", searchCmd.Name())
		fmt.Fprintf(os.Stderr, "  %s [options] <filename>\n", searchCmd.Name())
		fmt.Fprintf(os.Stderr, "Options:\n")
		searchCmd.PrintDefaults()
	}

	return SearchCommand{
		cmd:         searchCmd,
		root:        rootDir,
		insensitive: insensitive,
	}
}

func (command *SearchCommand) execute() {
	command.cmd.Parse(os.Args[2:])
	if command.cmd.NArg() < 1 {
		fmt.Println("expected filename")
		command.cmd.Usage()
		os.Exit(1)
	}

	command.fileName = command.cmd.Arg(0)

	filePaths, err := search.SearchFile(*command.root, command.fileName, *command.insensitive) 
	if err != nil {
		writeError("ERROR", "", err)
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
