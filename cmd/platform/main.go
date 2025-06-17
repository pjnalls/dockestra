package main

import (
	"os"

	"github.com/dockestra/command"
)

func main() {
	switch os.Args[1] {
	case "run":
		command.Run(os.Args[2:])
	default:
		panic("bad command")
	}
}
