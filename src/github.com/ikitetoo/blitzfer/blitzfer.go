package main

import "fmt"
import "flag"
//import "strings"
//import "io/ioutil"

var source string

func main() {
	flag.StringVar(&source, "directory", ".", "Path of directory to scan.")
	flag.Parse()
	fmt.Printf("%v\n", source)

//	if !source.contains("/", ".") {
//		fmt.Printf("Not a valid directory")
//	}

}
