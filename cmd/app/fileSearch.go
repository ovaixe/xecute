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

func searchCommand(cmd *flag.FlagSet, searchFileName *string) {

	cmd.Parse(os.Args[2:])
	if *searchFileName == "" {
		cmd.Usage()
		os.Exit(1)
	}

	fmt.Println("subcommand 'search'")
	fmt.Println("filename: ", *searchFileName)

	filePath, err := searchFile(*searchFileName)
	if err != nil {
		writeError("ERROR", "", err)
		os.Exit(1)
	}

	fmt.Println("File Path: ", filePath)
}

func searchFile(name string) (string, error) {
	root := "/"

	var filePath string

	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			writeError("Error accessing", path, err)
			return nil
		}

		if !info.IsDir() && info.Name() == name {
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
