package tests

import (
	"context"
	"encoding/json"
	"log"
	"testing"
	"vndb-go/vndb-go"
)

func TestVnQuery(t *testing.T) {
	if clientTest == nil {
		t.Logf("No client found")
		return
	}

	client := clientTest

	var vns []vndb_go.Vn
	vnQuery := vndb_go.NewVnQuery(client)
	vnQuery.Fields("id")
	vnQuery.Results(10)

	vnQuery.Filters(
		vndb_go.DevStatus.Equal("0"))

	vns, err := vnQuery.Get(context.TODO())
	if err != nil {
		t.Errorf("%s", err)
	}

	for _, vn := range vns {
		if vn.Id != nil {
			t.Logf("Found %s", *vn.Id)
		}
	}
}

func TestCharQuery(t *testing.T) {
	if clientTest == nil {
		t.Logf("No client found")
		return
	}

	client := clientTest
	charQuery := vndb_go.NewCharacterQuery(client)
	charQuery.Fields("id, description")
	charQuery.Results(10)
	charQuery.Page(1)

	log.Printf("Filters: %s", charQuery.BaseQuery.Query.Filters)

	var chars []vndb_go.Character
	chars, err := charQuery.Get(context.TODO())
	if err != nil {
		t.Errorf("%s", err)
	}

	for _, char := range chars {
		if char.Id != nil && char.Description != nil {
			t.Logf("ID: %s \nDesc: %s\n", *char.Id, *char.Description)
		} else {
			t.Logf("Character without description, ID: %s", *char.Id)
		}

		if char.Image != nil {
			t.Logf("ID: %s Image: %d", *char.Id, char.Image.Dims)
		}

		if char.Birthday != nil {
			t.Logf("Bday: %d", *char.Birthday)
		}
	}
}

func TestProducerQuery(t *testing.T) {
	if clientTest == nil {
		t.Logf("No client found")
		return
	}

	client := clientTest
	q := &vndb_go.Query{
		Page:    3,
		Results: 10,
		Fields:  "name",
	}

	ctx := context.TODO()
	results, err := client.Query(ctx, "producer", q)
	if err != nil {
		t.Fatal(err)
	}

	var prods []vndb_go.Producer
	if err := json.Unmarshal(results.Results, &prods); err != nil {
		t.Fatal(err)
	}

	for _, prod := range prods {
		if prod.Name != nil {
			t.Logf("Name: %s", *prod.Name)
		}
	}
}

func TestReleaseQuery(t *testing.T) {
	if clientTest == nil {
		t.Logf("No client found")
		return
	}

	client := clientTest
	q := &vndb_go.Query{
		Page:    1,
		Results: 50,
		Fields:  "id,producers.id",
	}

	ctx := context.TODO()
	results, err := client.Query(ctx, "release", q)
	if err != nil {
		t.Fatal(err)
	}

	var releases []vndb_go.Release
	if err := json.Unmarshal(results.Results, &releases); err != nil {
		t.Fatal(err)
	}

	for _, rl := range releases {
		if rl.Producers != nil {
			for _, prod := range *rl.Producers {
				if prod.Id != nil {
					t.Logf("Producer: %s", *prod.Id)
				}
			}
		}
	}
}
