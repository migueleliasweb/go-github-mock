package gen

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/go-kit/log"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/go-kit/log/level"
)

type ScrapeResult struct {
	HTTPMethod      string
	EndpointPattern string
}

// Replacing deprecated method strings.Title
// requires a "Caser"
var Caser = cases.Title(language.English)

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
// It can have an optional prefix, needed between method and endpoint
func FormatToGolangVarName(l log.Logger, sr ScrapeResult, pr string) string {
	result := Caser.String(strings.ToLower(sr.HTTPMethod))
	result += pr

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
			// avoid ending up with method like
			// PostEnterpriseEnterprisesActionsRunnerGroupsByEnterprise
			// by only adding part if different from the prefix
			if pr == "" || !strings.Contains(Caser.String(p), pr) {
				result = result + Caser.String(p)
			}

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
				result += Caser.String(splitPart)
			}
		}
	}

	return result
}

func FormatToGolangVarNameAndValue(l log.Logger, sr ScrapeResult, p string) string {
	sr = applyMutation(sr)

	return fmt.Sprintf(
		`var %s EndpointPattern = EndpointPattern{
	Pattern: "%s",
	Method:  "%s",
}
`,
		FormatToGolangVarName(l, sr, p),
		sr.EndpointPattern,
		strings.ToUpper(sr.HTTPMethod),
	) + "\n"
}
