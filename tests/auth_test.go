package tests

import (
	"context"
	"os"
	"testing"
	"vndb-go/vndb-go"
)

func TestAuth(t *testing.T) {
	client := vndb_go.NewVndbClient("")

	ctx := context.Background()
	auth, err := client.GetAuthInfo(ctx, os.Getenv("VNDB_TOKEN"))
	if err != nil {
		t.Errorf("Error getting auth info: %v", err)
	}

	t.Logf("auth info: \nId: %s \nUsername: %s \nPermissions: %v", auth.Id, auth.Username, auth.Permissions)
}
