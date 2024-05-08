package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

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
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	if filePath == "" {
		fmt.Println("File not found")
		os.Exit(1)
	}

	fmt.Println("File Path: ", filePath)
}

func searchFile(name string) (string, error) {
	root := "/"

	var filePath string

	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing %s: %v\n", path, err)
			return nil
		}

		if !info.IsDir() && info.Name() == name {
			filePath = path
		}

		return nil
	})

	if err != nil {
		return filePath, err
	}

	return filePath, nil
}
