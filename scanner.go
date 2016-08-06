package main

import (
        "fmt"
        "os"
	"log"
)

func scanDir(dir FsMetaData) {
	// Declare what we're scanning.
	if (debug == true) {
		log.Printf("[d] %v\n", dir.path)
	}

	// Open the directory for scanning.
	d, err := os.Open(dir.path)
	if err != nil {
		log.Printf("[%v]\n", err)
	}
	defer d.Close()

	// Start scanning or fault.
	f, err := d.Readdir(-1)
	for _, finfo := range f {
		path := dir.path+"/"+finfo.Name()

		m := finfo.Mode()

		// This should be switch/case
		if finfo.IsDir() {
			d := setDir(path, finfo)
//			pq <- d
			fmt.Printf("[d] %v\n", path)
			scanDir(d)
			escInsert(d)
			continue
		}

		if m.IsRegular() {
			setFile(path, finfo)
			fmt.Printf("[f] %v\n", path)
			continue
		}

                /*
                  I don't recall why I noted this... derp.
                  os.ModeSetuid      // u: setuid
                  os.ModeSetgid      // g: setgid
                  os.ModeSticky      // t: sticky
                */

		switch finfo.Mode() & os.ModeSymlink {

			case os.ModeSymlink:     // L: Symbolic Link
				rl, err := os.Readlink(path)
				if err != nil {
					log.Printf("%v\n", err)
				}
				setLink(path, finfo)
				fmt.Printf("[L] %v -> %v\n", path, rl)

			case os.ModeDevice:      // D: device
				setDevice(path, finfo)
				fmt.Printf("[D] %v\n", path)

			case os.ModeNamedPipe:   // p: named pipe (FIFO)
				setFifo(path, finfo)
				fmt.Printf("[p] %v\n", path)

			case os.ModeSocket:      // S: Unix domain socket
				setSocket(path, finfo)
				fmt.Printf("[S] %v\n", path)

			case os.ModeCharDevice:  // c: Unix character device, when ModeDevice is set
				setCharDev(path, finfo)
				fmt.Printf("[c] %v\n", path)

			default:
				fmt.Printf("[u] %v\n", path)
		}

	}
}

func scanInit(path string) {

        finfo, err := os.Stat(path)
        if err != nil {
	    log.Fatal(err)
        }

        if finfo.IsDir() {
	    d := setDir(path, finfo)
	    scanDir(d)
//          This must have been for debugging or something.
//	    var input string
//	    fmt.Scanln(&input)
        } else {
	    log.Fatal("["+path+"] is not a directory.")
	}

	return
}
