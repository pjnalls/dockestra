package main

import (
	"os"

	"github.com/pjnalls/dockestra/pkg/command"
)

func main() {
	switch os.Args[1] {
	case "run":
		command.Run(os.Args)
	default:
		panic("bad command")
	}
}
