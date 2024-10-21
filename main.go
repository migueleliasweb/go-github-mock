package main

import (
	"bytes"
	"flag"
	"log/slog"
	"os"
	"os/exec"

	"github.com/buger/jsonparser"

	"github.com/migueleliasweb/go-github-mock/src/gen"
)

var debug bool

func init() {
	flag.BoolVar(&debug, "debug", false, "output debug information")
}

func main() {
	flag.Parse()

	fetchAndWriteAPIDefinition()
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
		slog.Info("error executing gofmt", "err", err.Error())
		errorsFound = true
	}

	// to catch everything else (hopefully)
	if err := exec.Command("go", "vet", "./...").Run(); err != nil {
		slog.Info("error executing go vet", "err", err.Error())
		errorsFound = true
	}

	if errorsFound {
		os.Exit(1)
	}
}
