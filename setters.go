package main

import (
        "os"
        "path/filepath"
)

func setDir(path string, fi os.FileInfo) FsMetaData {
        var d FsMetaData
        d.path = path
        d.info = fi
        d.mode = fi.Mode()
        d.parent = filepath.Dir(path)
	d.ntype = "directory"
        return d
}

func setFile(path string, fi os.FileInfo) FsMetaData {
        var f FsMetaData
        f.path = path
        f.info = fi
        f.mode = fi.Mode()
        f.parent = filepath.Dir(path)
	f.ntype = "regularfile"
        return f
}

func setLink(path string, fi os.FileInfo) FsMetaData {
        var l FsMetaData
        l.path = path
        l.info = fi
        l.mode = fi.Mode()
        l.parent = filepath.Dir(path)
	l.ntype = "link"
        return l
}

func setDevice(path string, fi os.FileInfo) FsMetaData {
        var dev FsMetaData
        dev.path = path
        dev.info = fi
        dev.mode = fi.Mode()
        dev.parent = filepath.Dir(path)
	dev.ntype = "device"
        return dev
}

func setFifo(path string, fi os.FileInfo) FsMetaData {
        var f FsMetaData
        f.path = path
        f.info = fi
        f.mode = fi.Mode()
        f.parent = filepath.Dir(path)
	f.ntype = "fifo"
        return f
}

func setSocket(path string, fi os.FileInfo) FsMetaData {
        var f FsMetaData
        f.path = path
        f.info = fi
        f.mode = fi.Mode()
        f.parent = filepath.Dir(path)
	f.ntype = "socket"
        return f
}

func setCharDev(path string, fi os.FileInfo) FsMetaData {
        var f FsMetaData
        f.path = path
        f.info = fi
        f.mode = fi.Mode()
        f.parent = filepath.Dir(path)
	f.ntype = "chardevice"
        return f
}
