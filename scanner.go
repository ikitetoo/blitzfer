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

func SetLink(path string, fi os.FileInfo) FsMetaData {
	var l FsMetaData
	l.path = path
	l.info = fi
	l.mode = fi.Mode()
	l.parent = filepath.Dir(path)
	return l
}

func ScanDir(dir FsMetaData) {
	if DEBUG {
//		fmt.Printf("[d] ", dir.path)
		fmt.Printf("%v\n", dir.path)
	}

	d, err := os.Open(dir.path)
	if err != nil {
		log.Printf("%v", err)
//		log.Fatal(err)
	}
	defer d.Close()

	f, err := d.Readdir(-1)
	for _, finfo := range f {
		path := dir.path+"/"+finfo.Name()

		m := finfo.Mode()

		// This should be switch/case
		if finfo.IsDir() {
			d := SetDir(path, finfo)
			ScanDir(d)
		}

		if m.IsRegular() {
			SetFile(path, finfo)
//			fmt.Printf("[f] %v\n", path)
			fmt.Printf("%v\n", path)
		}

		if finfo.Mode() & os.ModeSymlink == os.ModeSymlink {
			SetLink(path, finfo)
			fmt.Printf("[l] ", path)
			fmt.Printf("%v\n", path)
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
//                fmt.Printf("Root Dir Found: %v\n", path)
            }

	    d := SetDir(path, finfo)
	    ScanDir(d)
//	    var input string
//	    fmt.Scanln(&input)
            return

        } else {

	    log.Fatal("["+path+"] is not a directory.")

	}
}
