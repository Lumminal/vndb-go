# vndb-go

vndb-go is an api wrapper built with Golang as a learning project,
therefore it may be unstable for actual use.

Zero-dependencies are required to run this library.

Look at the examples below to see how it works.

# Installation

Run
```go install https://github.com/Lumminal/vndb-go```

# Examples

### Creating a client and grabbing the site's stats

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
	// There's many filters you can choose from
    
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

GPL-3.0