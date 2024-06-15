package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var ErrFileNotFound = errors.New("FILE NOT FOUND")

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

	err := command.walkFiles()
	if err != nil {
		writeError("ERROR", "", err)
		os.Exit(1)
	}

	command.printResults()
}

func (command *SearchCommand) walkFiles() error {
	var filePaths []string

	err := filepath.Walk(*command.root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		switch {
		case *command.insensitive:
			if !info.IsDir() && strings.Contains(strings.ToLower(info.Name()), strings.ToLower(command.fileName)) {
				filePaths = append(filePaths, path)
			}
		default:
			if !info.IsDir() && info.Name() == command.fileName {
				filePaths = append(filePaths, path)
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	if len(filePaths) == 0 {
		return ErrFileNotFound
	}

	command.result = filePaths
	return nil
}

func (command *SearchCommand) printResults() {
	fmt.Println("Results:")
	for _, res := range command.result {
		fmt.Print("File Path: ")
		fmt.Println(res)
	}
}
