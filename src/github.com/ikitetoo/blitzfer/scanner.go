package main

import (
        "fmt"
        "os"
        "path/filepath"
	"log"
)

func ScanDir(dir FsMetaData) {
	fmt.Printf("%v\n", dir.path)

	d, err := os.Open(dir.path)
	if err != nil {
		fmt.Printf("here\n")
		log.Fatal(err)
	}

	f, err := d.Readdir(-1)
	for _, fi := range f {
		fmt.Println(fi.Name())
	}
}

func ScanInit(path string) {
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
                fmt.Printf("Root Dir Found: %v\n", path)
            }

            d.path = path
            d.info = finfo
            d.mode = mode
            d.parent = filepath.Dir(path)

            if DEBUG {
                fmt.Printf("Dir Struct: [%v]\n", d)
            }

	    ScanDir(d)
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
