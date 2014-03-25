package main

import (
	"fmt"
	"flag"
)

var DEBUG bool

const (
        VERSION = "0.0.1"
)

func main() {
	var source_path string
	DEBUG = true

	flag.BoolVar(&DEBUG, "verbose", false, "Verbose output")
	flag.StringVar(&source_path, "directory", ".", "Path of directory to scan.")
	flag.Parse()

	fmt.Printf("Blitzfer Version: %v\n", VERSION)
	ScanInit(source_path)
}
