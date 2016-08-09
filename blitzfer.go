package main

import (
	"os"
	"log"
	"fmt"
	"flag"
	"sync"
        "github.com/olivere/elastic"
)

// Main Variables
var debug bool
var ver   bool
var max   int
var esc   *elastic.Client
var sourcePath string
var configFile string

// Passive Queue channel
var pq chan FsMetaData = make(chan FsMetaData)

// Active queue channel
var aq chan FsMetaData = make(chan FsMetaData)

// Allow us to wait on routines to finish.
var done sync.WaitGroup

// Create a uid Dictionary, to limit lookups.
var uidToNameMap map[uint32]string
var gidToNameMap map[uint32]string

const (
        version = "0.1.0"
)

var myConfigs *Config

func main() {
        // Allocate map structures.
        uidToNameMap = make(map[uint32]string)
        gidToNameMap = make(map[uint32]string)

        // Add default config path.
        flag.StringVar(&configFile, "configFile", "./config.yml", "Path of config file.")
	flag.BoolVar(&debug, "verbose", false, "Verbose output")
	flag.StringVar(&sourcePath, "directory", ".", "Path of directory to scan.")
	flag.IntVar(&max, "max", 100, "Max number of concurently open directories.")
	flag.BoolVar(&ver, "version", false, "Version output")
	flag.Parse()

	if (ver == true) {
		fmt.Printf("Blitzfer Version: %v\n", version)
		return
	}

	// listen for new directories.
        go passiveQueue()

	// load configuration map.
	myConfigs = loadConfig()

	os.Exit(0)

	// test elasticsearch connection.
	esc = escConnect()

	// Start directory scanning.
	scanInit(sourcePath)

	log.Printf("Done.\n")
	os.Exit(0)
}
