package gen

var enabledMutators = map[string]func(ScrapeResult) ScrapeResult{
	"/repos/{owner}/{repo}/contents/{path}": func(sr ScrapeResult) ScrapeResult {
		sr.EndpointPattern = "/repos/{owner}/{repo}/contents/{path:.+}"

		return sr
	},
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
