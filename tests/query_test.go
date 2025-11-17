package tests

import (
	"context"
	"encoding/json"
	"os"
	"testing"
	"vndb-go/wrapper"

	"github.com/joho/godotenv"
)

func TestVnQuery(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Logf("Error loading .env file")
	}

	client := wrapper.NewVndbClient(os.Getenv(wrapper.VNDBToken))
	q := &wrapper.Query{
		Page:    1,
		Results: 100,
		Fields:  "id, image.url",
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
	}
}
