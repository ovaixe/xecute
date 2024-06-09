package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

var ErrFileNotFound = errors.New("FILE NOT FOUND")

type SearchCommand struct {
	cmd      *flag.FlagSet
	root     *string
	fileName string
}

func NewSearchCommand() SearchCommand {
	searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
	rootDir := searchCmd.String("root", "/", "Root directory")

	searchCmd.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", searchCmd.Name())
		fmt.Fprintf(os.Stderr, "  %s [options] <filename>\n", searchCmd.Name())
		fmt.Fprintf(os.Stderr, "Options:\n")
		searchCmd.PrintDefaults()
	}

	return SearchCommand{
		cmd:  searchCmd,
		root: rootDir,
	}
}

func (command SearchCommand) execute() {

	command.cmd.Parse(os.Args[2:])
	if command.cmd.NArg() < 1 {
		fmt.Println("expected filename")
		command.cmd.Usage()
		os.Exit(1)
	}

	command.fileName = command.cmd.Arg(0)

	fmt.Println("subcommand 'search'")
	fmt.Println("filename: ", command.fileName)

	filePath, err := command.searchFile()
	if err != nil {
		writeError("ERROR", "", err)
		os.Exit(1)
	}

	fmt.Println("File Path: ", filePath)
}

func (command SearchCommand) searchFile() (string, error) {
	var filePath string

	err := filepath.Walk(*command.root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			writeError("Error accessing", path, err)
			return nil
		}

		if !info.IsDir() && info.Name() == command.fileName {
			filePath = path
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	if filePath == "" {
		return "", ErrFileNotFound
	}

	return filePath, nil
}
