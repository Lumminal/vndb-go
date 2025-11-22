package main

import (
	"context"
	"log"
	"os"

	"github.com/Lumminal/vndb-go"
)

func main() {
	client := vndb_go.NewVndbClient(os.Getenv(vndb_go.VNDBToken))

	var vns []vndb_go.Vn                  // we will store the results here
	vnQuery := vndb_go.NewVnQuery(client) // create a new query
	vnQuery.Fields("title", "devstatus")  // grab only vns with "title" and "devstatus"
	vnQuery.Results(10)                   // return only 10 VNs

	// OPTIONAL:
	// You can add filters like:
	// vnQuery.Filters(vndb_go.DevStatus.Equal("0"))

	vns, err := vnQuery.Get(context.TODO()) // get the query results
	if err != nil {
		log.Fatal(err)
	}

	// print them out to the console
	for _, vn := range vns {
		if vn.Id != nil { // make sure it exists before accessing it
			log.Printf("Found %s", *vn.Id)
		}
		if vn.DevStatus != nil {
			log.Printf("DevStatus %d", *vn.DevStatus)
		}
	}
}
