package tests

import (
	"context"
	"os"
	"testing"

	"github.com/Lumminal/vndb-go"
)

func TestAuth(t *testing.T) {
	client := vndb_go.NewVndbClient("")

	expected := &vndb_go.AuthInfo{
		BaseUser: vndb_go.BaseUser{
			Id:       "u227260",
			Username: "Luminal",
		},
		Permissions: []string{},
	}

	ctx := context.Background()
	auth, err := client.GetAuthInfo(ctx, os.Getenv("VNDB_TOKEN"))
	if err != nil {
		t.Errorf("Error getting auth info: %s", err)
	}

	if !vndb_go.CompareAuthInfo(expected, auth) {
		t.Errorf("Got %v, but expected %v", auth, expected)
	}
}
