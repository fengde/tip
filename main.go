package main

import (
	"os"
)

func main() {
	if len(os.Args) > 1 {
		Show(os.Args[1])
		return
	}
	Show()
}
