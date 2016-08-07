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
	// TODO: This needs to go into a config file... preferably .yml
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

        // Determine node ownerships.
	nodeUid := node.info.Sys().(*syscall.Stat_t).Uid
	nodeGid := node.info.Sys().(*syscall.Stat_t).Gid

        // Need to convert these for LookupId, derp.
        nodeUidString := fmt.Sprint(nodeUid)
        nodeGidString := fmt.Sprint(nodeGid)

        // Test to see if the Owner UID already exists in memory/map. Otherwise add it.
	_, nodeUidExists := uidToNameMap[nodeUid]
	if (nodeUidExists) {
		fmt.Printf("[%v] exists in map\n", nodeUid)
	} else {
		// Insert uid[username] into map.
		u, _ := user.LookupId(nodeUidString)
		uidToNameMap[nodeUid] = u.Username
		fmt.Printf("will insert [%v / %v] into our user map\n", nodeUid, u.Username)
	}

        // Rinse, repeat for GID table.
	_, nodeGidExists := gidToNameMap[nodeGid]
	if (nodeGidExists) {
		fmt.Printf("[%v] exists in map\n", nodeGid)
	} else {
		// Insert gid[gidname] into map.
		g, _ := user.LookupId(nodeGidString)
		gidToNameMap[nodeGid] = g.Username
		fmt.Printf("will insert [%v / %v] into our gid map\n", nodeGid, g.Username)
	}

        // Shove this into Elasticsearch.
	type esDoc struct {
		ModTime         time.Time
		IsDir           bool
		ParentDirectory	string
		AbsPath         string
                OwnerId         uint32
                GidId           uint32
                Perms           os.FileMode
                OwnerName       string
                GidName         string
	}

        // Create out Elasticsearch document payload.
        doc := esDoc {
		ModTime:         node.info.ModTime(),
		IsDir:           node.info.IsDir(),
		ParentDirectory: node.parent,
                AbsPath:         node.path,
		OwnerId:         nodeUid,
		GidId:           nodeGid,
		Perms:           node.mode,
		OwnerName:       uidToNameMap[nodeUid],
		GidName:         gidToNameMap[nodeGid],
	}

	// Marshal the data prior to sending it up to Elastic Search.
        b, err := json.Marshal(doc)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("Elastic Search Payload: %s\n", b)
}
