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

func SetDevice(path string, fi os.FileInfo) FsMetaData {
	var d FsMetaData
	d.path = path
	d.info = fi
	d.mode = fi.Mode()
	d.parent = filepath.Dir(path)
	return d
}

func ScanDir(dir FsMetaData) {
	if DEBUG {
		fmt.Printf("[d] %v\n", dir.path)
	}

	d, err := os.Open(dir.path)
	if err != nil {
		log.Printf("%v", err)
	}
	defer d.Close()

	f, err := d.Readdir(-1)
	for _, finfo := range f {
		path := dir.path+"/"+finfo.Name()

		m := finfo.Mode()

		// This should be switch/case
		if finfo.IsDir() {
			d := SetDir(path, finfo)
			SetDir(path, finfo)
			fmt.Print("[d] %v\n", path)
			ScanDir(d)
			continue
		}

		if m.IsRegular() {
			SetFile(path, finfo)
			fmt.Printf("[f] %v\n", path)
			continue
		}
/*
os.ModeSetuid      // u: setuid
os.ModeSetgid      // g: setgid
os.ModeSticky      // t: sticky
*/

		switch finfo.Mode() & os.ModeSymlink {

			case os.ModeSymlink:     // L: Symbolic Link
				rl, err := ReadLink(path)
				if err != nil {
					log.Printf("%v", err)
				}
				SetLink(path, finfo)
				fmt.Printf("[L] %v\n", path)

			case os.ModeDevice:      // D: device
				SetDevice(path, finfo)
				fmt.Printf("[D] %v\n", path)

/*
			case os.ModeNamedPipe:   // p: named pipe (FIFO)
				SetFifo(path, finfo)
				fmt.Printf("[p] %v\n", path)

			case os.ModeSocket:      // S: Unix domain socket
				SetSocket(path, finfo)
				fmt.Printf("[S] %v\n", path)

			case os.ModeCharDevice:  // c: Unix character device, when ModeDevice is set
				SetCharDev(path, finfo)
				fmt.Printf("[c] %v\n", path)
*/

			default:
				fmt.Printf("[u] %v\n", path)
		}

	}
}

func ScanInit(path string) {

        finfo, err := os.Stat(path)
        if err != nil {
	    log.Fatal(err)
        }

        if finfo.IsDir() {

	    d := SetDir(path, finfo)
	    ScanDir(d)
//	    var input string
//	    fmt.Scanln(&input)
            return

        } else {

	    log.Fatal("["+path+"] is not a directory.")

	}
}
