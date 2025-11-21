package tests

import (
	"context"
	"os"
	"testing"
	"vndb-go/vndb-go"
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

	if !Compare(expected, auth) {
		t.Errorf("Got %v, but expected %v", auth, expected)
	}
}

func Compare(a, b *vndb_go.AuthInfo) bool {
	if a.Username != b.Username || a.Id != b.Id {
		return false
	}

	if len(a.Permissions) != len(b.Permissions) {
		return false
	}

	for i := range a.Permissions {
		if a.Permissions[i] != b.Permissions[i] {
			return false
		}
	}

	return true
}
