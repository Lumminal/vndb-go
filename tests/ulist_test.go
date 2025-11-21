package tests

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/Lumminal/vndb-go"
)

func TestUlist(t *testing.T) {
	client := clientTest

	user := "u227260"

	ulist, err := client.GetUListLabels(context.TODO(), &user)
	if err != nil {
		t.Fatal(err)
	}

	for _, label := range ulist.Labels {
		log.Printf("%s", label.Label)
	}
}

func TestUlistPost(t *testing.T) {
	client := clientTest

	q := vndb_go.NewUlistQuery(client)
	q.SetUser("u227260")
	q.Fields("lastmod")

	results, err := q.Get(context.TODO())
	if err != nil {
		t.Fatal(err)
	}

	for _, result := range results {
		log.Printf("%v", time.Unix(result.LastMod, 0))
	}
}
