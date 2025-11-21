package tests

import (
	"context"
	"os"
	"testing"

	"github.com/Lumminal/vndb-go"
	"github.com/Lumminal/vndb-go/types"
)

func TestAuth(t *testing.T) {
	client := vndb_go.NewVndbClient("")

	expected := &types.AuthInfo{
		BaseUser: types.BaseUser{
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

	if !types.CompareAuthInfo(expected, auth) {
		t.Errorf("Got %v, but expected %v", auth, expected)
	}
}
