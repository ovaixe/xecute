package main

import (
	"fmt"
)

const RED = "\033[31m"
const RESET = "\033[0m"

func writeError(title, info string, err error) {
	fmt.Print(RED + title + "\t" + RESET)
	fmt.Println(info, err)
}
