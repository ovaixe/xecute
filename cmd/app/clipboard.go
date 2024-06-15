package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
)

type ClipboardCommand struct {
	cmd      *flag.FlagSet
	dir      *string
	fileName string
}

func NewClipboardCommand() ClipboardCommand {
	clipboardCmd := flag.NewFlagSet("xclip", flag.ExitOnError)
	clipboardDir := clipboardCmd.String("dir", ".", "Directory")

	clipboardCmd.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", clipboardCmd.Name())
		fmt.Fprintf(os.Stderr, "  %s [options] <filename>\n", clipboardCmd.Name())
		fmt.Fprintf(os.Stderr, "Options:\n")
		clipboardCmd.PrintDefaults()
	}

	return ClipboardCommand{
		cmd: clipboardCmd,
		dir: clipboardDir,
	}
}

func (command ClipboardCommand) execute() {
	command.cmd.Parse(os.Args[2:])

	if command.cmd.NArg() < 1 {
		fmt.Println("expected filename")
		command.cmd.Usage()
		os.Exit(0)
	}

	command.fileName = command.cmd.Arg(0)

	err := command.writeAll()
	if err != nil {
		writeError("ERROR", "", err)
		os.Exit(1)
	}
}

func (command ClipboardCommand) writeAll() error {
	file := *command.dir + "/" + command.fileName

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
