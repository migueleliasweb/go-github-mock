package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/buger/jsonparser"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/migueleliasweb/go-github-mock/src/gen"
)

var debug bool

func init() {
	flag.BoolVar(&debug, "debug", false, "output debug information")
}

func main() {
	flag.Parse()

	var l log.Logger

	l = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))

	l = log.With(l, "caller", log.DefaultCaller)

	if debug {
		l = level.NewFilter(l, level.AllowDebug())
		level.Debug(l).Log("msg", "running in debug mode")
	} else {
		l = level.NewFilter(l, level.AllowInfo())
	}

	fetchAndWriteAPIDefinition(l)
}

// type helper just to ensure uniqueness of the generated output
type uniq map[string]struct{}

func (u uniq) has(s string) bool {
	_, ok := u[s]
	return ok
}

func fetchAndWriteAPIDefinition(l log.Logger) {

	buf := bytes.NewBuffer([]byte(gen.OUTPUT_FILE_HEADER))
	u := make(uniq)
	defs := [][]byte{
		gen.FetchAPIDefinition(l, gen.GITHUB_OPENAPI_DEFINITION_LOCATION),
		gen.FetchAPIDefinition(l, gen.GITHUB_OPENAPI_ENTERPRISE_DEFINITION_LOCATION),
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
						l,
						gen.ScrapeResult{
							HTTPMethod:      httpMethod,
							EndpointPattern: endpointPattern,
						},
					)
					if !u.has(code) {
						u[code] = struct{}{}
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
		level.Error(l).Log("msg", fmt.Sprintf("error executing gofmt: %s", err.Error()))
		errorsFound = true
	}

	// to catch everything else (hopefully)
	if err := exec.Command("go", "vet", "./...").Run(); err != nil {
		level.Error(l).Log("msg", fmt.Sprintf("error executing go vet: %s", err.Error()))
		errorsFound = true
	}

	if errorsFound {
		os.Exit(1)
	}
}
