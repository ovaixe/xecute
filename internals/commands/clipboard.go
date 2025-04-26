package commands

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ovaixe/xecute/internals/clipboard"
	"github.com/ovaixe/xecute/internals/utils"
)

type ClipboardCommand struct {
	CMD      *flag.FlagSet
	dir      *string
	fileName string
}

func NewClipboardCommand() ClipboardCommand {
	clipboardCmd := flag.NewFlagSet("c", flag.ExitOnError)
	clipboardDir := clipboardCmd.String("d", "./", "Directory")

	clipboardCmd.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s: copies file content to clipboard\n", clipboardCmd.Name())
		fmt.Fprintf(os.Stderr, "  %s [options] <filename>\n", clipboardCmd.Name())
		fmt.Fprintf(os.Stderr, "Options:\n")
		clipboardCmd.PrintDefaults()
	}

	return ClipboardCommand{
		CMD: clipboardCmd,
		dir: clipboardDir,
	}
}

func (command *ClipboardCommand) Execute() {
	command.CMD.Parse(os.Args[2:])

	if command.CMD.NArg() < 1 {
		fmt.Println("expected filename")
		command.CMD.Usage()
		os.Exit(0)
	}

	command.fileName = command.CMD.Arg(0)
	dir := *command.dir
	if !strings.HasSuffix(dir, "/") {
		dir += "/"
	}

	file := dir + command.fileName

	data, err := os.ReadFile(file)
	if err != nil {
		utils.WriteError("ERROR", "", err)
		os.Exit(1)
	}

	err = clipboard.Write(data)
	if err != nil {
		utils.WriteError("ERROR", "", err)
		os.Exit(1)
	}

	fmt.Println("Text copied to clipboard")
}
