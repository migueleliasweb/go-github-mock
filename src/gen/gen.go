package gen

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

const GITHUB_OPENAPI_DEFINITION_LOCATION = "https://github.com/github/rest-api-description/blob/main/descriptions/api.github.com/api.github.com.json?raw=true"

const OUTPUT_FILE_HEADER = `package mock

// Code generated; DO NOT EDIT.

`

const OUTPUT_FILEPATH = "src/mock/endpointpattern.go"

type ScrapeResult struct {
	HTTPMethod      string
	EndpointPattern string
}

func FetchAPIDefinition() ([]byte, error) {
	resp, err := http.Get(GITHUB_OPENAPI_DEFINITION_LOCATION)
	if err != nil {
		return nil, fmt.Errorf("error fetching github's api definition: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading github's api definition: %w", err)
	}
	return bodyBytes, nil
}

// FormatToGolangVarName generated the proper golang variable name
// given a endpoint format from the API
func FormatToGolangVarName(sr ScrapeResult) string {
	result := strings.Title(strings.ToLower(sr.HTTPMethod))

	if sr.EndpointPattern == "/" {
		return result + "Slash"
	}

	// handles urls with dashes in them
	pattern := strings.ReplaceAll(sr.EndpointPattern, "-", "/")

	// cleans up varname when pattern was mutated
	// e.g see `GetReposContentsByOwnerByRepoByPath`
	re := regexp.MustCompile(`[a-zA-Z0-9\/\{\}\_]+`)
	matches := re.FindAllString(pattern, -1)
	pattern = strings.Join(matches, "")

	epSplit := strings.Split(
		pattern,
		"/",
	)

	// handle the first part of the variable name
	for _, part := range epSplit {
		if len(part) < 1 || string(part[0]) == "{" {
			continue
		}

		splitPart := strings.Split(part, "_")

		for _, p := range splitPart {
			result = result + strings.Title(p)
		}
	}

	// handle the "By`X`" part of the variable name
	for _, part := range epSplit {
		if len(part) < 1 {
			continue
		}

		if string(part[0]) == "{" {
			part = strings.ReplaceAll(part, "{", "")
			part = strings.ReplaceAll(part, "}", "")

			result += "By"

			for _, splitPart := range strings.Split(part, "_") {
				result += strings.Title(splitPart)
			}
		}
	}

	return result
}

func FormatToGolangVarNameAndValue(sr ScrapeResult) string {
	sr = applyMutation(sr)

	return fmt.Sprintf(
		`var %s EndpointPattern = EndpointPattern{
	Pattern: "%s",
	Method:  "%s",
}
`,
		FormatToGolangVarName(sr),
		sr.EndpointPattern,
		strings.ToUpper(sr.HTTPMethod),
	) + "\n"
}
