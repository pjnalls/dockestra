package command

import "fmt"

func Run(params []string) {
	fmt.Printf("Running %v\n", params[2:])
}
