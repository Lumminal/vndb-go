package tests

import (
	"context"
	"testing"
)

func TestUser(t *testing.T) {
	client := clientTest

	const badUser = "NoUserWithThisNameExists"
	const goodUser = "yorhel"

	ctx := context.Background()
	user, _ := client.GetUser(badUser, ctx)

	if user != nil {
		t.Errorf("expecting nil user but got non-nill user")
	}

	usr2, err := client.GetUser(goodUser, ctx)
	if err != nil {
		t.Fatal("expecting non-nil user but got nill user")
	}

	if usr2.UserFields.LengthVotes != 0 || usr2.UserFields.LengthVotesSum != 0 {
		t.Fatalf("expecting 0 votes but got %d and sum %d", usr2.UserFields.LengthVotes, usr2.UserFields.LengthVotesSum)
	}

	usr3, _ := client.GetUserWithFields(goodUser, ctx, true, false)

	if usr3.LengthVotes == 0 {
		t.Fatalf("expecting non-zero votes but got %d", usr3.LengthVotes)
	}
	if usr3.LengthVotesSum != 0 {
		t.Fatalf("expecting 0 vote sum but got %d", usr3.LengthVotesSum)
	}
}
