package main

import (
	"fmt"
	"os"
)

func main() {
	// get current path
	if currentPath, err := os.Getwd(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	} else {

	}
}
