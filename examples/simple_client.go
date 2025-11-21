package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"vndb-go"
)

func main() {
	client := vndb_go.NewVndbClient(os.Getenv(vndb_go.VNDBToken)) // make a client via "VNDB_TOKEN" env variable

	ctx := context.Background()
	stats, err := client.GetStats(ctx) // grab stats
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Visual Novels: %d", stats.Vn)
}
