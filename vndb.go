package main

import (
	"fmt"
	"os"
	"vndb-go/vndb-go"
)

func main() {
	fmt.Println(os.Getenv(vndb_go.VNDBToken))
}
