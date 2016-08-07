package main
import ( "fmt"
	 "os"
	 "os/user"
	 "log"
	 "time"
	 "syscall"
	 "encoding/json"
         "github.com/olivere/elastic" )

func escConnect() {
	// Create a client and connect to http://192.168.2.10:9201
        log.Printf("Connecting to http://192.168.99.100:9200 ...\n")
        esc, err := elastic.NewClient(
                    elastic.SetURL("http://192.168.99.100:9200"),
		    elastic.SetSniff(false),
                    elastic.SetMaxRetries(10))
	if err != nil {
		// Handle error
		log.Printf("Error: %v\n", err)
		log.Fatal("Failed to connect to elasticsearch.")
	} else {
		if (debug == true) {
			fmt.Printf("Got connection: [%v]\n", esc)
		}

		// Use the IndexExists service to check if a specified index exists.
		exists, err := esc.IndexExists("files").Do()
		if err != nil {
		    // Handle error
		    panic(err)
		}

		if !exists {
		    // Create a new index.
		    createIndex, err := esc.CreateIndex("files").Do()
		    if err != nil {
		        // Handle error
		        panic(err)
		    }
		    if !createIndex.Acknowledged {
		        // Not acknowledged
			log.Printf("Failed to create Index: [%v]", err)
		    }
		}
	}
}

func escUpdate(node FsMetaData) {

        // Shove this into Elasticsearch
	type esDoc struct {
		modTime         time.Time
		isDir           bool
		parentDirectory	string
		absPath         string
//              ownerName       string
//              gidName         string
                ownerId         uint32
                gidId           uint32
                perms           os.FileMode
	}

        doc := esDoc {
		modTime:         node.info.ModTime(),
		isDir:           node.info.IsDir(),
		parentDirectory: node.parent,
                absPath:         node.path,
		ownerId:         node.info.Sys().(*syscall.Stat_t).Uid,
		gidId:           node.info.Sys().(*syscall.Stat_t).Gid,
		perms:           node.mode,
	}

        b, err := json.Marshal(doc)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("-----> %v", b)
}
