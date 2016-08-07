package main

import (
        "os"
)

type FsMetaData struct {
        path string
        info os.FileInfo
	mode os.FileMode
        parent string
	ntype string
}
