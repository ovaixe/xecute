package main

import (
	"flag"
	"fmt"
	"os"
)

var version = "1.0.0"

var buildTime string

var CMDS = []string{"search"}

func main() {
	searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
	searchFileName := searchCmd.String("filename", "", "Search file name")

	displayVersion := flag.Bool("version", false, "Display version and Build time")

	flag.Parse()

	if *displayVersion {
		fmt.Printf("Version:\t%s\n", version)
		fmt.Printf("Build time:\t%s\n", buildTime)
		os.Exit(0)
	}

	if len(os.Args) < 2 {
		flag.Usage()
		fmt.Println("Usage: xecute <subcommand> [options]")
		fmt.Println("Available subcommands:")
		fmt.Println(CMDS)
		searchCmd.Usage()
		os.Exit(0)
	}

	switch os.Args[1] {
	case "search":
		searchCommand(searchCmd, searchFileName)
	default:
		fmt.Println("expected subcommand")
		fmt.Println("Available commands:")
		fmt.Println(CMDS)
		os.Exit(1)
	}

}
