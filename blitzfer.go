package main

import (
	"fmt"
	"flag"
	"sync"
)

var debug bool
var ver   bool
var max   int

// Passive Queue channel
var pq chan FsMetaData = make(chan FsMetaData)

// Active queue channel
var aq chan FsMetaData = make(chan FsMetaData)

// Allow us to wait on routines to finish.
var done sync.WaitGroup

const (
        version = "0.0.1"
)

func main() {
	var source_path string

	debug = true

	flag.BoolVar(&debug, "verbose", false, "Verbose output")
	flag.StringVar(&source_path, "directory", ".", "Path of directory to scan.")
	flag.IntVar(&max, "max", 100, "Max number of concurently open directories.")
	flag.BoolVar(&ver, "version", false, "Version output")
	flag.Parse()

	if (ver == true) {
		fmt.Printf("Blitzfer Version: %v\n", version)
		return
	}

	// listen for new directories.
        go passiveQueue()

	// Start directory scanning.
	scanInit(source_path)
}
