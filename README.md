# vndb-go

vndb-go is an api wrapper for [VNDB API](https://api.vndb.org/kana).
Zero dependencies are required to run this package.

Check out the examples to see how it works.

### Status 

This project was made as part of learning backend development and golang,
therefore **it's experimental and unstable for actual use**.

Some features like UList PATCH/DELETE have not been implemented.

# Installation

Run
```go get github.com/Lumminal/vndb-go```

# Examples

### Creating a client and grabbing general VNDB stats

```go
package main

import (
	"context"
	"fmt"
	"log"
	"vndb-go/vndb-go"
)

func main() {
	client := vndb_go.NewVndbClient(YOUR_TOKEN) // you can also pass an empty string
	
    // note: you can also set an environment variable "VNDB_TOKEN" and access it like this:
	// client := vndb_go.NewVndbClient(os.Getenv(vndb_go.VNDBToken))
	
	ctx := context.TODO()
	stats, err := client.GetStats(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Characters: %d", stats.Chars)
}

```
 
### Running a query on visual novels

```go
package main

import (
	"vndb-go/vndb-go"
	"context"
	"fmt"
	"log"
)

func main() {

	client := vndb_go.NewVndbClient(YOUR_TOKEN)
	
    var vns []vndb_go.Vn // we will store the results here
    vnQuery := vndb_go.NewVnQuery(client) // create a new query
    vnQuery.Fields("title", "devstatus") // grab only vns with "title" and "devstatus"
    vnQuery.Results(10) // return only 10 VNs
    
    // OPTIONAL: 
	// You can add filters like:
	// vnQuery.Filters(vndb_go.DevStatus.Equal("0"))
	// which will give you only results that are
	// DevStatus is equal to 0 (devstatus = 0)
	// There's many filters you can choose from.
    
    vns, err := vnQuery.Get(context.TODO()) // get the query results
    if err != nil {
		log.Fatal(err)
    }

    for _, vn := range vns {
        if vn.Id != nil {
            log.Printf("Found %s", *vn.Id)
        }
		if vn.DevStatus != nil {
			log.Printf("DevStatus %s", *vn.DevStatus)
        }
    }
} 
```

# ULists

Ulists have not been implemented fully, but they support POST/GET methods.
They have their own query and you have to specify the user's id before making
a query by calling:

```yourUlistQuery.SetUser(USER_ID_HERE)```

# Issues

If you find any problems with the code, please open an issue and I'm gonna try
to fix it.

# [TODO]

[] Add better tests

[] Add patching/delete for ulists

[] Properly test all features

# License

The code is licensed under GPL-3.0