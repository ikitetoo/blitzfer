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
var ver        bool
var max        int
var esc        *elastic.Client
var configFile string
var debug      bool
var sourcePath string
var esIp       string
var esPort     string
var esIndex    string

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

// Pointers to respective configs
var conf *BlitzferConfigs

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

        // Display version
	if (ver == true) {
	  fmt.Printf("Blitzfer Version: %v\n", version)
          return
	}

	// listen for new directories. Will likely abandon this idea... let leave it for now.
        go passiveQueue()

	// load configuration map.
        conf       = loadBlitzferConfigs()
        debug      = conf.Configs["configs"].Blitzfer.Debug
        sourcePath = conf.Configs["configs"].Blitzfer.Directory
        esIp       = conf.Configs["configs"].Elasticsearch.Ip
        esPort     = conf.Configs["configs"].Elasticsearch.Port
        esIndex    = conf.Configs["configs"].Elasticsearch.Index

        if ( debug == true ) {
           fmt.Printf("\n---------- Settings ---------\n")
           fmt.Printf("%-20v %v\n", "Debug:", debug)
           fmt.Printf("%-20v %v\n", "Directory:", sourcePath)
           fmt.Printf("%-20v %v\n", "Elasticsearch IP:", esIp)
           fmt.Printf("%-20v %v\n", "Elasticsearch Port:", esPort)
           fmt.Printf("%-20v %v\n", "Elasticsearch Index:", esIndex)
           fmt.Printf("-----------------------------\n\n")
        }

	// Setup elasticsearch connection.
	esc = escConnect()

	os.Exit(0)

	// Start directory scanning.
	scanInit(sourcePath)

	log.Printf("Done.\n")
	os.Exit(0)
}
