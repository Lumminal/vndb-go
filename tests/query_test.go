package tests

import (
	"context"
	"encoding/json"
	"testing"
	"vndb-go/wrapper"
)

func TestVnQuery(t *testing.T) {
	if clientTest == nil {
		t.Logf("No client found")
		return
	}

	client := clientTest
	q := &wrapper.Query{
		Page:    1,
		Results: 10,
		Fields:  "id, image.url, released",
		Reverse: true,
	}

	ctx := context.TODO()
	results, err := client.Query(ctx, "vn", q)
	if err != nil {
		t.Fatal(err)
	}

	var vns []wrapper.Vn
	if err := json.Unmarshal(results.Results, &vns); err != nil {
		t.Fatal(err)
	}

	for _, vn := range vns {
		if vn.Image != nil {
			t.Logf("ID: %s Image: %s", *vn.Id, *vn.Image.Url)
		}

		if vn.Released != nil {
			t.Logf("ID: %s Released: %s", *vn.Id, vn.Released.Release)
		}
	}
}

func TestCharQuery(t *testing.T) {
	if clientTest == nil {
		t.Logf("No client found")
		return
	}

	client := clientTest
	q := &wrapper.Query{
		Page:    1,
		Results: 10,
		Fields:  "id, description, image.dims",
	}

	ctx := context.TODO()
	results, err := client.Query(ctx, "character", q)
	if err != nil {
		t.Fatal(err)
	}

	var chars []wrapper.Character
	if err := json.Unmarshal(results.Results, &chars); err != nil {
		t.Fatal(err)
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
	}
}

func TestProducerQuery(t *testing.T) {
	if clientTest == nil {
		t.Logf("No client found")
		return
	}

	client := clientTest
	q := &wrapper.Query{
		Page:    3,
		Results: 10,
		Fields:  "name",
	}

	ctx := context.TODO()
	results, err := client.Query(ctx, "producer", q)
	if err != nil {
		t.Fatal(err)
	}

	var prods []wrapper.Producer
	if err := json.Unmarshal(results.Results, &prods); err != nil {
		t.Fatal(err)
	}

	for _, prod := range prods {
		if prod.Name != nil {
			t.Logf("Name: %s", *prod.Name)
		}
	}
}
