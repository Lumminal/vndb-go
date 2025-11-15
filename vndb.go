package main

import (
	"context"
	"fmt"
	"log"
	"vndb-go/wrapper"
)

func main() {
	c := wrapper.NewVndbClient("9bry-bu11z-bqy87-aao3-z8qk8-e8jx5-a6o1")

	stats, err := c.GetUserWithFields("yorhel", context.TODO(), true, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("User's id is %d", stats.LengthVotes)
}
