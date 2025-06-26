package mock

import (
	"testing"

	"github.com/google/go-github/v73/github"
	"github.com/gorilla/mux"
)

func TestWithRequestMatchEnterprise(t *testing.T) {
	option := WithRequestMatchEnterprise(
		GetApp,
		github.App{
			Name: func() *string {
				name := "myapp"

				return &name
			}(),
		},
	)

	router := mux.NewRouter()

	option(router)

	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathRegexp, err := route.GetPathRegexp()

		if err != nil {
			t.Fatalf("got error reading path regexp: %s", err)
		}

		if pathRegexp != "^/api/v3/app$" {
			t.Errorf("pathRegexp is %s, want ^/api/v3/app$", pathRegexp)
		}

		return nil
	})
}

func TestWithRequestMatchPagesEnterprise(t *testing.T) {
	option := WithRequestMatchPagesEnterprise(
		GetOrgsReposByOrg,
		[]github.Repository{
			{
				Name: github.Ptr("repo-A-on-first-page"),
			},
			{
				Name: github.Ptr("repo-B-on-first-page"),
			},
		},
	)

	router := mux.NewRouter()

	option(router)

	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathRegexp, err := route.GetPathRegexp()

		if err != nil {
			t.Fatalf("got error reading path regexp: %s", err)
		}

		expectedRegexp := "^/api/v3/orgs/(?P<v0>[^/]+)/repos$"

		if pathRegexp != expectedRegexp {
			t.Errorf("pathRegexp is %s, want %s", pathRegexp, expectedRegexp)
		}

		return nil
	})
}
