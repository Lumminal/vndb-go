package tests

import (
	"context"
	"testing"
	"vndb-go/wrapper"
)

func TestAuth(t *testing.T) {
	client := wrapper.NewVndbClient("")

	ctx := context.Background()
	auth, err := client.GetAuthInfo(ctx, "9bry-bu11z-bqy87-aao3-z8qk8-e8jx5-a6o1")
	if err != nil {
		t.Errorf("Error getting auth info: %v", err)
	}

	t.Logf("auth info: \nId: %s \nUsername: %s \nPermissions: %v", auth.Id, auth.Username, auth.Permissions)
}
