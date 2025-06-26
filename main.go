package main

import (
	"bytes"
	"context"
	"io/fs"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/buger/jsonparser"
	"github.com/google/go-github/v73/github"

	"github.com/migueleliasweb/go-github-mock/src/gen"
	"golang.org/x/mod/modfile"
)

func main() {
	fetchAndWriteAPIDefinition()
	updateGoGithubDep()
}

// type helper just to ensure uniqueness of the generated output
type uniq map[string]struct{}

// has will add `s` to the map whe not found
func (u uniq) has(s string) bool {
	_, has := u[s]

	if !has {
		u[s] = struct{}{}
	}

	return has
}

func fetchAndWriteAPIDefinition() {
	buf := bytes.NewBuffer([]byte(gen.OUTPUT_FILE_HEADER))

	u := make(uniq)

	defs := [][]byte{
		gen.FetchAPIDefinition(gen.GITHUB_OPENAPI_DEFINITION_LOCATION),
		gen.FetchAPIDefinition(gen.GITHUB_OPENAPI_ENTERPRISE_DEFINITION_LOCATION),
	}

	for _, d := range defs {
		jsonparser.ObjectEach(
			d,
			func(key, endpointDefinition []byte, _ jsonparser.ValueType, _ int) error {
				endpointPattern := string(key)

				httpMethods := []string{}

				jsonparser.ObjectEach(
					endpointDefinition,
					func(key, _ []byte, _ jsonparser.ValueType, _ int) error {
						httpMethods = append(httpMethods, string(key))

						return nil
					},
				)

				for _, httpMethod := range httpMethods {
					code := gen.FormatToGolangVarNameAndValue(
						gen.ScrapeResult{
							HTTPMethod:      httpMethod,
							EndpointPattern: endpointPattern,
						},
					)
					if !u.has(code) {
						buf.WriteString(code)
					}
				}

				return nil
			},
			"paths",
		)
	}

	os.WriteFile(
		gen.OUTPUT_FILEPATH,
		buf.Bytes(),
		0755,
	)

	errorsFound := false

	// to catch possible format errors
	if err := exec.Command("gofmt", "-w", gen.OUTPUT_FILEPATH).Run(); err != nil {
		slog.Error("error executing gofmt", "err", err.Error())
		errorsFound = true
	}

	// to catch everything else (hopefully)
	if err := exec.Command("go", "vet", "./...").Run(); err != nil {
		slog.Error("error executing go vet", "err", err.Error())
		errorsFound = true
	}

	if errorsFound {
		os.Exit(1)
	}
}

func updateGoGithubDep() {
	ghClient := github.NewClient(nil)

	releaseInfo, _, err := ghClient.Repositories.GetLatestRelease(
		context.Background(),
		"google",
		"go-github",
	)

	if err != nil {
		slog.Error(
			"error fetching latest release from google/go-github",
			"err", err.Error(),
		)

		os.Exit(1)
	}

	goGithubMockFileBytes, err := os.ReadFile("go.mod")

	if err != nil {
		slog.Error(
			"error reading go.mod contents from go-github-mock",
			"err", err.Error(),
		)

		os.Exit(1)
	}

	goGithubMockModFile, err := modfile.Parse(
		"go.mod",
		goGithubMockFileBytes,
		nil,
	)

	if err != nil {
		slog.Error(
			"error parsing go.mod contents from go-github-mock",
			"err", err.Error(),
		)

		os.Exit(1)
	}

	localGoGithubPath := ""
	localGoGithubVersion := ""

	for _, requireInfo := range goGithubMockModFile.Require {
		if strings.HasPrefix(requireInfo.Mod.Path, "github.com/google/go-github/v") {
			localGoGithubPath = requireInfo.Mod.Path
			localGoGithubVersion = requireInfo.Mod.Version
			break
		}
	}

	// e.g. "v71.0.0" => "v69"
	latestGoGithubPath := "github.com/google/go-github/" + strings.Split(*releaseInfo.TagName, ".")[0]
	latestGoGithubVersion := *releaseInfo.TagName

	// if versions are the same, exit early
	if localGoGithubPath == latestGoGithubPath && localGoGithubVersion == latestGoGithubVersion {
		slog.Info(
			"go-github dependency already on latest release, skipping upgrade",
			"local-go-github", localGoGithubPath,
			"local-go-github-version", localGoGithubVersion,
			"latest-go-github", latestGoGithubPath,
			"latest-go-github-version", latestGoGithubVersion,
		)
		return
	}

	if localGoGithubPath == "" || localGoGithubVersion == "" {
		slog.Error("unable to find local go-github version information")
		os.Exit(1)
	}

	filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if strings.HasSuffix(path, ".go") {
			fileBytes, err := os.ReadFile(path)

			if err != nil {
				return err
			}

			fileContentsStr := string(fileBytes)

			fileContentsStr = strings.ReplaceAll(
				fileContentsStr,
				localGoGithubPath,
				latestGoGithubPath,
			)

			fileContentsStr = strings.ReplaceAll(
				fileContentsStr,
				localGoGithubVersion,
				latestGoGithubVersion,
			)

			if err := os.WriteFile(
				path,
				[]byte(fileContentsStr),
				os.FileMode(os.O_TRUNC),
			); err != nil {
				slog.Error(
					"failed to write update to file",
					"file", path,
					"err", err.Error(),
				)

				os.Exit(1)
			}
		}

		return nil
	})

	// to catch possible format errors
	if err := exec.Command("gofmt", "-w", gen.OUTPUT_FILEPATH).Run(); err != nil {
		slog.Info("error executing gofmt", "err", err.Error())
	}
}
