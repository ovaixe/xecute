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
	fileName *string
}

func NewSearchCommand() SearchCommand {
	searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
	searchFileName := searchCmd.String("filename", "", "Search file name")

	return SearchCommand{
		cmd:      searchCmd,
		fileName: searchFileName,
	}
}

func (command SearchCommand) execute() {

	command.cmd.Parse(os.Args[2:])
	if *command.fileName == "" {
		command.cmd.Usage()
		os.Exit(1)
	}

	fmt.Println("subcommand 'search'")
	fmt.Println("filename: ", *command.fileName)

	filePath, err := command.searchFile()
	if err != nil {
		writeError("ERROR", "", err)
		os.Exit(1)
	}

	fmt.Println("File Path: ", filePath)
}

func (command SearchCommand) searchFile() (string, error) {
	root := "/"

	var filePath string

	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			writeError("Error accessing", path, err)
			return nil
		}

		if !info.IsDir() && info.Name() == *command.fileName {
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
