package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
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
		return errors.New("FAILED TO READ FILE")
	}

	copyCmd := exec.Command("xclip", "-selection", "clipboard")
	copyCmd.Stdin = io.NopCloser(bytes.NewReader(data))
	err = copyCmd.Run()
	if err != nil {
		return errors.New("FAILED TO COPY TO CLIPBOARD: ")
	}

	fmt.Println("Text copied to clipboard successfully!")
	return nil
}
