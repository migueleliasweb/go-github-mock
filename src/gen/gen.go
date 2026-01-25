package gen

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const GITHUB_OPENAPI_DEFINITION_LOCATION = "https://github.com/github/rest-api-description/blob/main/descriptions/api.github.com/api.github.com.json?raw=true"

const GITHUB_OPENAPI_ENTERPRISE_DEFINITION_LOCATION = "https://raw.githubusercontent.com/github/rest-api-description/main/descriptions/ghec/ghec.json"

const OUTPUT_FILE_HEADER = `package mock

// Code generated; DO NOT EDIT.

`
const OUTPUT_FILEPATH = "src/mock/endpointpattern.go"

// Title is a Caser that replaces deprecated method strings.Title.
// requires a "Caser".
var Title = cases.Title(language.English)

type ScrapeResult struct {
	HTTPMethod      string
	EndpointPattern string
}

type Exit struct{ Code int }

func exitHandler() {
	if e := recover(); e != nil {
		if exit, ok := e.(Exit); ok {
			os.Exit(exit.Code)
		}

		panic(e) // re-panic if not our Exit
	}
}

func FetchAPIDefinition(d string) []byte {
	defer exitHandler()

	resp, err := http.Get(d)
	if err != nil {
		slog.Info(
			"error fetching github's api definition",
			"err", err.Error(),
		)

		panic(Exit{1}) // exit with code 1, cleanup runs
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Info(
			"error fetching github's api definition",
			"err", err.Error(),
		)

		panic(Exit{1}) // exit with code 1, cleanup runs
	}

	return bodyBytes
}

// FormatToGolangVarName generated the proper golang variable name
// given a endpoint format from the API.
func FormatToGolangVarName(sr ScrapeResult) string {
	result := Title.String(strings.ToLower(sr.HTTPMethod))

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

		splitPart := strings.SplitSeq(part, "_")

		for p := range splitPart {
			result += Title.String(p)
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

			for splitPart := range strings.SplitSeq(part, "_") {
				result += Title.String(splitPart)
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
