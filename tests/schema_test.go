package tests

import (
	"context"
	"testing"
	"vndb-go/wrapper"
)

func TestSchema(t *testing.T) {
	client := wrapper.NewVndbClient("9bry-bu11z-bqy87-aao3-z8qk8-e8jx5-a6o1")

	ctx := context.TODO()
	schema, err := client.GetSchema(ctx)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Schema: %v", schema)

	// test producer
	for _, extlink := range schema.ExtLinks.Producer {
		if extlink.Id == nil {
			continue
		}

		t.Logf("ExtLinks: %v", *extlink.Url)
	}
}
