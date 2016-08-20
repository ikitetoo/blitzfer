package main
import ( "fmt"
	 "os"
	 "os/user"
	 "log"
	 "time"
	 "syscall"
         "github.com/olivere/elastic" )

func escConnect() *elastic.Client {

        if ( debug == true ) {
          log.Printf("Connecting to http://%v:%v ...\n", esIp, esPort)
        }

	var esUrl string
        esUrl = fmt.Sprintf("http://%s:%s", esIp, esPort)

        esc, err := elastic.NewClient(
                    elastic.SetURL(esUrl),
		    elastic.SetSniff(false),
                    elastic.SetMaxRetries(10)) // TODO: Add more ES options in yml for stuff like this.

	if err != nil {
		// Handle error
		log.Printf("Error: %v\n", err)
		log.Fatal("Failed to connect to Elasticsearch.")
	} else {
		if (debug == true) {
			fmt.Printf("Got connection: [%v]\n", esc)
		}

		// Use the IndexExists service to check if a specified index exists.
		exists, err := esc.IndexExists(esIndex).Do()
		if err != nil {
		    // Handle error
		    panic(err)
		}

		if !exists {
		    // Create a new index.
		    createIndex, err := esc.CreateIndex(esIndex).Do()
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

	return esc
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

	        // Waiting for go version 1.7 for this function... in the meantime.. hack city
		// https://tip.golang.org/pkg/os/user/#LookupGroup	
		// g, err := user.LookupGroupId(nodeGidString)

		g, err := user.LookupId(nodeGidString)
		if err != nil {
	            // Handle error
		    fmt.Printf("Node: %s", node.path)
		    gidToNameMap[nodeGid] = "unknown"
		    // panic(err)
		} else {
		    gidToNameMap[nodeGid] = g.Username
		}

		fmt.Printf("will insert [%s / %s] into our gid map\n", nodeGid, gidToNameMap[nodeGid])
	}

        // Shove this into Elasticsearch.
	type esDoc struct {
		MTIME           time.Time
		IsDir           bool
		PARENT          string
		PATH            string
		SIZE            int64
                UID             uint32
                GID             uint32
                MODE            os.FileMode
                USER            string
                GROUP           string
	}

        // Create out Elasticsearch document payload.
        esPayload := esDoc {
		MTIME:           node.info.ModTime(),
		IsDir:           node.info.IsDir(),
		PARENT:          node.parent,
                PATH:            node.path,
		SIZE:            node.info.Size(),
		UID:             nodeUid,
		GID:             nodeGid,
		MODE:            node.mode,
		USER:            uidToNameMap[nodeUid],
		GROUP:           gidToNameMap[nodeGid],
	}

	if (debug == true) {
		fmt.Printf("Elastic Search Payload: %s\n", esPayload)
	}

	// Send esPayload
	put1, err := esc.Index().Index(esIndex).Type(node.ntype).BodyJson(esPayload).Do()
	if err != nil {
	    // Handle error
	    panic(err)
	}

	fmt.Printf("Indexed file data %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)

	// Flush to make sure the documents got written.
/*	_, err = esc.Flush().Index(esIndex).Do()
	if err != nil {
	    fmt.Printf("E1\n")
	    panic(err)
	}
*/

}
