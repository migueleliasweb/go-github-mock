package mock

import (
	"context"
	"testing"

	"github.com/google/go-github/v37/github"
)

func TestNewMockedHttpClient(t *testing.T) {
	mockedHttpClient := NewMockedHttpClient(
		WithRequestMatch(
			GetUser,
			[][]byte{
				MustMarshal(github.User{
					Name: github.String("foobar"),
				}),
			},
		),
	)
	c := github.NewClient(mockedHttpClient)

	ctx := context.Background()

	user, _, userErr := c.Users.Get(ctx, "someUser")

	if user == nil || user.Name == nil || *user.Name != "foobar" {
		t.Fatalf("User name is %s, want foobar", user)
	}

	if userErr != nil {
		t.Errorf("User err is %s, want nil", userErr.Error())
	}

	orgs, _, err := c.Organizations.List(
		ctx,
		*(user.Name),
		nil,
	)

	if len(orgs) != 1 {
		t.Errorf("Orgs len is %d want 1", len(orgs))
	}

	if err != nil {
		t.Errorf("Err is %s, want nil", err.Error())
	}

	if *(orgs[0].Name) != "foobar123thisorgwasmocked" {
		t.Errorf("orgs[0].Name is %s, want %s", *orgs[0].Name, "foobar123thisorgdoesnotexist")
	}
}
