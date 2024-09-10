package gen

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/go-kit/log"

	"github.com/go-kit/log/level"
)

const GITHUB_OPENAPI_DEFINITION_LOCATION = "https://github.com/github/rest-api-description/blob/main/descriptions/api.github.com/api.github.com.json?raw=true"

const GITHUB_OPENAPI_ENTERPRISE_DEFINITION_LOCATION = "https://raw.githubusercontent.com/github/rest-api-description/main/descriptions/ghec/ghec.json"

const OUTPUT_FILE_HEADER = `package mock

// Code generated; DO NOT EDIT.

`
const OUTPUT_FILEPATH = "src/mock/endpointpattern.go"

type ScrapeResult struct {
	HTTPMethod      string
	EndpointPattern string
}

func FetchAPIDefinition(l log.Logger, d string) []byte {
	resp, err := http.Get(d)

	if err != nil {
		level.Error(l).Log(
			"msg", "error fetching github's api definition",
			"err", err.Error(),
		)

		os.Exit(1)
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		level.Error(l).Log(
			"msg", "error fetching github's api definition",
			"err", err.Error(),
		)

		os.Exit(1)
	}

	return bodyBytes
}

// FormatToGolangVarName generated the proper golang variable name
// given a endpoint format from the API
func FormatToGolangVarName(l log.Logger, sr ScrapeResult) string {
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

	//handle the "By`X`" part of the variable name
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

func FormatToGolangVarNameAndValue(l log.Logger, sr ScrapeResult) string {
	sr = applyMutation(sr)

	return fmt.Sprintf(
		`var %s EndpointPattern = EndpointPattern{
	Pattern: "%s",
	Method:  "%s",
}
`,
		FormatToGolangVarName(l, sr),
		sr.EndpointPattern,
		strings.ToUpper(sr.HTTPMethod),
	) + "\n"
}
