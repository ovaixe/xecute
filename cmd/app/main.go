package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ovaixe/xecute/internals/commands"
)

var version = "dev"

var buildTime string

var CMDS = []string{"s", "c"}

func main() {
	displayVersion := flag.Bool("version", false, "Display version and Build time")

	flag.Parse()

	if *displayVersion {
		fmt.Printf("Version:\t%s\n", version)
		fmt.Printf("Build time:\t%s\n", buildTime)
		os.Exit(0)
	}

	searchCmd := commands.NewSearchCommand()
	clipboardCmd := commands.NewClipboardCommand()

	if len(os.Args) < 2 {
		flag.Usage()
		fmt.Println("Usage: xecute <subcommand> [options]")
		fmt.Fprintln(os.Stdout, "Available subcommands: ", CMDS)
		searchCmd.CMD.Usage()
		clipboardCmd.CMD.Usage()
		os.Exit(0)
	}

	switch os.Args[1] {
	case "s":
		searchCmd.Execute()
	case "c":
		clipboardCmd.Execute()
	default:
		fmt.Println("expected subcommand")
		fmt.Println("Available commands:")
		fmt.Println(CMDS)
		os.Exit(1)
	}
}
