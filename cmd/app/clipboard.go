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
	dir      *string
	fileName *string
}

func NewClipboardCommand() ClipboardCommand {
	clipboardCmd := flag.NewFlagSet("clipboard", flag.ExitOnError)
	clipboardFilePath := clipboardCmd.String("filepath", "", "Text file path")
	clipboardDir := clipboardCmd.String("dir", "./", "Directory")
	clipboardFileName := clipboardCmd.String("filename", "", "Text file name")

	return ClipboardCommand{
		cmd:      clipboardCmd,
		filePath: clipboardFilePath,
		dir:      clipboardDir,
		fileName: clipboardFileName,
	}
}

func (command ClipboardCommand) execute() {
	command.cmd.Parse(os.Args[2:])
	if *command.filePath == "" && *command.fileName == "" {
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
	var file string

	if *command.filePath != "" {
		file = *command.filePath
	} else {
		file = *command.dir + "/" + *command.fileName
	}

	data, err := os.ReadFile(file)
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
