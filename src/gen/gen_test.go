package gen

import (
	"testing"

	"github.com/go-kit/log"
)

func TestFormatToGolangVarName(t *testing.T) {
	tests := []struct {
		name string
		sr   ScrapeResult
		want string
	}{
		{
			name: "simple",
			sr: ScrapeResult{
				EndpointPattern: "/repos/{owner}/{repo}/actions/artifacts",
				HTTPMethod:      "GET",
			},
			want: "GetReposActionsArtifactsByOwnerByRepo",
		},
		{
			name: "withParamUnderscore",
			sr: ScrapeResult{
				EndpointPattern: "/repos/{owner}/{repo}/actions/artifacts/{artifact_id}",
				HTTPMethod:      "GET",
			},
			want: "GetReposActionsArtifactsByOwnerByRepoByArtifactId",
		},
		{
			name: "withMultipleParamUnderscore",
			sr: ScrapeResult{
				EndpointPattern: "/users/{username}/packages/{package_type}/{package_name}",
				HTTPMethod:      "GET",
			},
			want: "GetUsersPackagesByUsernameByPackageTypeByPackageName",
		},
		{
			name: "withUrlWithDash",
			sr: ScrapeResult{
				EndpointPattern: "/orgs/{org}/actions/permissions/selected-actions",
				HTTPMethod:      "GET",
			},
			want: "GetOrgsActionsPermissionsSelectedActionsByOrg",
		},
		{
			name: "withUrlWithUnderscore",
			sr: ScrapeResult{
				EndpointPattern: "/repos/{owner}/{repo}/actions/runs/{run_id}/pending_deployments",
				HTTPMethod:      "GET",
			},
			want: "GetReposActionsRunsPendingDeploymentsByOwnerByRepoByRunId",
		},
		{
			name: "withUrlWithNumber",
			sr: ScrapeResult{
				EndpointPattern: "/repos/{owner}/{repo}/actions/runs/{run_id}/pending_deployments",
				HTTPMethod:      "GET",
			},
			want: "GetReposActionsRunsPendingDeploymentsByOwnerByRepoByRunId",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatToGolangVarName(log.NewNopLogger(), tt.sr); got != tt.want {
				t.Errorf("formatToGolangVarName() = %v, want %v", got, tt.want)
			}
		})
	}
}
