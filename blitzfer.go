package main

import (
	"fmt"
	"flag"
)

var DEBUG bool
var ver bool

const (
        VERSION = "0.0.1"
)

func main() {
	var source_path string
	DEBUG = true

	flag.BoolVar(&DEBUG, "verbose", false, "Verbose output")
	flag.StringVar(&source_path, "directory", ".", "Path of directory to scan.")
	flag.BoolVar(&ver, "version", false, "Version output")
	flag.Parse()

	if (ver == true) {
		fmt.Printf("Blitzfer Version: %v\n", VERSION)
		return
	}

	ScanInit(source_path)
}
