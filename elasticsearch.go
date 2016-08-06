package main
import ( "fmt"
	 "os"
         "github.com/olivere/elastic" )

func es_connect() {
	// Create a client and connect to http://192.168.2.10:9201
        fmt.Printf("Connecting to http://192.168.99.100:9200 ...\n")
        esc, err := elastic.NewClient(
                    elastic.SetURL("http://192.168.99.100:9200"),
		    elastic.SetSniff(false),
                    elastic.SetMaxRetries(10))
	if err != nil {
		// Handle error
                fmt.Printf("Failed to connect to elasticsearch!\n")
		fmt.Printf("Derp: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("Got connection: [%v]\n", esc)
	}
}
