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

var (
	ErrFileRead      = errors.New("FAILED TO READ FILE")
	ErrClipboardCopy = errors.New("FAILED TO COPY TO CLIPBOARD")
)

type ClipboardCommand struct {
	cmd      *flag.FlagSet
	filePath *string
}

func NewClipboardCommand() ClipboardCommand {
	clipboardCmd := flag.NewFlagSet("clipboard", flag.ExitOnError)
	clipboardFilePath := clipboardCmd.String("filepath", "", "Text file path")

	return ClipboardCommand{
		cmd:      clipboardCmd,
		filePath: clipboardFilePath,
	}
}

func (command ClipboardCommand) execute() {
	command.cmd.Parse(os.Args[2:])
	if *command.filePath == "" {
		command.cmd.Usage()
		os.Exit(1)
	}

	err := command.writeAll()
	if err != nil {
		writeError("ERROR", "", err)
		os.Exit(1)
	}
}

func (command ClipboardCommand) writeAll() error {
	data, err := os.ReadFile(*command.filePath)
	if err != nil {
		return err
	}

	copyCmd := exec.Command("xclip", "-selection", "clipboard")
	copyCmd.Stdin = io.NopCloser(bytes.NewReader(data))
	err = copyCmd.Run()
	if err != nil {
		return err
	}

	fmt.Println("Text copied to clipboard successfully!")
	return nil
}
