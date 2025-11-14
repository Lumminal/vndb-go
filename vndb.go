package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"vndb-go/wrapper"
)

func main() {
	c := http.Client{
		Timeout: 3 * time.Second,
	}

	stats, err := wrapper.GetStats(&c, context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("There's %d characters and %d Visual Novels", stats.Chars, stats.Vn)
}
