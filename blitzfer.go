package main

import "fmt"
//import "io/ioutil"
import "flag"

var source string

func main() {
	flag.StringVar(&source, "directory", ".", "Path of directory to scan.")
	flag.Parse()
	fmt.Printf("%v\n", source)
}
