package main

import (
        "fmt"
        "os"
        "path/filepath"
	"log"
)

func SetDir(path string, fi os.FileInfo) FsMetaData {
	var d FsMetaData
	d.path = path
	d.info = fi
	d.mode = fi.Mode()
	d.parent = filepath.Dir(path)
	return d
}

func SetFile(path string, fi os.FileInfo) FsMetaData {
	var f FsMetaData
	f.path = path
	f.info = fi
	f.mode = fi.Mode()
	f.parent = filepath.Dir(path)
	return f
}

func ScanDir(dir FsMetaData) {
	if DEBUG {
		fmt.Printf("Scanning: %v\n", dir.path)
	}

	d, err := os.Open(dir.path)
	if err != nil {
		log.Fatal(err)
	}

	f, err := d.Readdir(-1)
	for _, finfo := range f {
		path := dir.path+"/"+finfo.Name()

		m := finfo.Mode()
		if finfo.IsDir() {
			d := SetDir(path, finfo)
			ScanDir(d)
		}

		if m.IsRegular() {
			SetFile(path, finfo)
			fmt.Printf("File: %v\n", path)
		}
	}
}

func ScanInit(path string) {

        finfo, err := os.Stat(path)
        if err != nil {
	    log.Fatal(err)
        }

        if finfo.IsDir() {

            if DEBUG {
                fmt.Printf("Root Dir Found: %v\n", path)
            }

	    d := SetDir(path, finfo)
	    ScanDir(d)
            return

        } else {

	    log.Fatal("["+path+"] is not a directory.")

	}
}
