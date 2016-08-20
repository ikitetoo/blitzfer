package main

import (
	"path/filepath"
	"io/ioutil"
        "gopkg.in/yaml.v2"
        "github.com/davecgh/go-spew/spew"
)

// Yaml Structures for Defaults and Elasticsearch configs

type BlitzferConfigs struct {
    Configs map[string] Options
}

// Config yaml structure.
type Options struct {

    Blitzfer struct {
        Debug     bool
        Directory string
    }

    Elasticsearch struct {
        Ip    string
        Port  string
        Index string
    }

}

func loadBlitzferConfigs() *BlitzferConfigs {

    // Read our config file
    filename, _ := filepath.Abs(configFile)
    yamlFile, err := ioutil.ReadFile(filename)

    if err != nil {
        panic(err)
    }

    var blitzferConfigs BlitzferConfigs

    // UnMarshal the config into our config structure.
    err = yaml.Unmarshal(yamlFile, &blitzferConfigs.Configs)
    if err != nil {
        panic(err)
    }

    if ( blitzferConfigs.Configs["configs"].Blitzfer.Debug == true ) {
      spew.Dump(blitzferConfigs.Configs)
    }

    return &blitzferConfigs
}
