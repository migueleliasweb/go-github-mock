package gen

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	OptionalLastParameter = "{%s:.*}"
	ExtendedLastParameter = "{%s:.+}"
)

var enabledMutations = map[string]string{
	"/repos/{owner}/{repo}/contents/{path}":                     OptionalLastParameter,
	"/repos/{owner}/{repo}/git/ref/{ref}":                       ExtendedLastParameter,
	"/repos/{owner}/{repo}/git/refs/{ref}":                      ExtendedLastParameter,
	"/repos/{owner}/{repo}/commits/{ref}":                       ExtendedLastParameter,
	"/repos/{owner}/{repo}/issues/{issue_number}/labels/{name}": ExtendedLastParameter,
	"/orgs/{org}/actions/runners/{runner_id}/labels/{name}":     ExtendedLastParameter,
	"/repos/{owner}/{repo}/labels/{name}":                       ExtendedLastParameter,
}

func mutator(formatString string) func(ScrapeResult) ScrapeResult {
	return func(sr ScrapeResult) ScrapeResult {
		endpointSplits := strings.Split(sr.EndpointPattern, "/")
		lastParam := endpointSplits[len(endpointSplits)-1]

		r := regexp.MustCompile(`[a-z]+`)

		lastParamCleaned := r.FindString(lastParam)

		endpointSplits[len(endpointSplits)-1] = fmt.Sprintf(formatString, lastParamCleaned)

		sr.EndpointPattern = strings.Join(endpointSplits, "/")

		return sr
	}
}

// applyMutation applies mutation to the scrape result if necessary.
//
// There are some edge cases due to inconsistencies between GitHub's OpenAPI definition
// compared to the real world.
func applyMutation(sr ScrapeResult) ScrapeResult {
	if mutation, found := enabledMutations[sr.EndpointPattern]; found {
		return mutator(mutation)(sr)
	}

	return sr
}
