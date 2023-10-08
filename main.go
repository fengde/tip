package main

import (
	"os"
	"tip/cmd"
)

func main() {
	if len(os.Args) > 1 {
		cmd.Show(os.Args[1])
		return
	}
	cmd.Show()
}
