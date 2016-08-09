package main

import (
        "fmt"
	"path/filepath"
	"io/ioutil"
        "gopkg.in/yaml.v2"
)

// Defaults and Elasticsearch configs
type Config struct {
    Configs map[string]Options
    Elasticsearch map[string]Elasticsearch
}

type Options struct {
    Debug     bool
    Directory string
}

type Elasticsearch struct {
    Index     string
    Ip        string
    Port      int
}

func loadConfig() *Config {

    // Read our config file
    filename, _ := filepath.Abs(configFile)
    yamlFile, err := ioutil.ReadFile(filename)

    if err != nil {
        panic(err)
    }

    var config Config

    // UnMarshal the config into our config structure.
    err = yaml.Unmarshal(yamlFile, &config)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Value: %#v\n", config.Configs)

    return &config
}
