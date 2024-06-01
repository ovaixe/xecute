package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func copyCammand(cmd *flag.FlagSet, filepath *string) {
	cmd.Parse(os.Args[2:])
	if *filepath == "" {
		cmd.Usage()
		os.Exit(1)
	}

	err := writeAll(*filepath)
	if err != nil {
		writeError("ERROR", "", err)
		os.Exit(1)
	}
}

func writeAll(filepath string) error {
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
		return errors.New("Failed to read file")
	}

	copyCmd := exec.Command("xclip", "-selection", "clipboard")
	copyCmd.Stdin = io.NopCloser(bytes.NewReader(data))
	err = copyCmd.Run()
	if err != nil {
		log.Fatalf("Failed to copy to clipboard: %v", err)
		return errors.New("Failed to copy to clipboard")
	}

	fmt.Println("Text copied to clipboard successfully!")
	return nil
}
