package gen

import (
	"fmt"
	"regexp"
	"strings"
)

var enabledMutators = map[string]func(ScrapeResult) ScrapeResult{
	"/repos/{owner}/{repo}/contents/{path}":                     allowExtendedLastParamMutatorHelper(),
	"/repos/{owner}/{repo}/git/ref/{ref}":                       allowExtendedLastParamMutatorHelper(),
	"/repos/{owner}/{repo}/git/refs/{ref}":                      allowExtendedLastParamMutatorHelper(), // thanks for the consistency, GitHub
	"/repos/{owner}/{repo}/commits/{ref}":                       allowExtendedLastParamMutatorHelper(), // thanks for not base64encode the parameter, GitHub
	"/repos/{owner}/{repo}/issues/{issue_number}/labels/{name}": allowExtendedLastParamMutatorHelper(),
}

// allowExtendedLastParamMutatorHelper mutates the last param of the endpoint pattern
// allowing it to have any characters (including slashes)
func allowExtendedLastParamMutatorHelper() func(ScrapeResult) ScrapeResult {
	return func(sr ScrapeResult) ScrapeResult {
		endpointSplits := strings.Split(sr.EndpointPattern, "/")
		lastParam := endpointSplits[len(endpointSplits)-1]

		r := regexp.MustCompile(`[a-z]+`)

		lastParamCleaned := r.FindString(lastParam)

		endpointSplits[len(endpointSplits)-1] = fmt.Sprintf("{%s:.+}", lastParamCleaned)

		sr.EndpointPattern = strings.Join(endpointSplits, "/")

		return sr
	}
}

// applyMutation applies mutation to the scrape result if necessary.
//
// There are some edge cases due to inconsistencies between GitHub's OpenAPI definition
// compared to the real world.
func applyMutation(sr ScrapeResult) ScrapeResult {
	if mutator, found := enabledMutators[sr.EndpointPattern]; found {
		return mutator(sr)
	}

	return sr
}
