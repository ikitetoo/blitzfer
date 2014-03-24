package main

import (
	"fmt"
	"flag"
	"os"
	"path/filepath"
)

var DEBUG bool
const (
        VERSION = "0.0.1"
)

func main() {
	var source_path string
	DEBUG = true

	flag.StringVar(&source_path, "directory", ".", "Path of directory to scan.")
	flag.Parse()

	fmt.Printf("Blitzfer Version: %v\n", VERSION)
	scandir(source_path)
}

func scandir (path string) {
	var d FsMetaData
	var f FsMetaData

	finfo, err := os.Stat(path)
	if err != nil {
	    fmt.Printf("%v no such file or directory\n", path)
	}
	mode := finfo.Mode()

	// Directories
	if finfo.IsDir() {
	    if DEBUG {
		fmt.Printf("Dir Found: %v\n", path)
		return
            }
	    d.path = path
	    d.info = finfo
	    d.mode = mode
	    d.parent = filepath.Dir(path)
	    if DEBUG {
	    	fmt.Printf("Dir Struct: [%v]\n", d)
	    }

	    return
	}

	// Files
	if mode.IsRegular() {
            if DEBUG {	
   	        fmt.Printf("File Found: %v\n", path)
		return
	    }
	    f.path = path
	    f.info = finfo
	    f.mode = mode
	    f.parent = filepath.Dir(path)
	    if DEBUG {
		fmt.Printf("File Struct: [%v]\n", f)
	    }
	
	    return
	}
	
}
