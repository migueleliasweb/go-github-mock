package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
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

	apiDefinition, err := gen.FetchAPIDefinition()
	if err != nil {
		log.Fatal(err)
	}

	buf := bytes.NewBuffer([]byte(gen.OUTPUT_FILE_HEADER))

	jsonparser.ObjectEach(
		apiDefinition,
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

				buf.WriteString(code)
			}

			return nil
		},
		"paths",
	)

	ioutil.WriteFile(
		gen.OUTPUT_FILEPATH,
		buf.Bytes(),
		0o755,
	)

	// to catch possible format errors
	if err := exec.Command("gofmt", "-w", gen.OUTPUT_FILEPATH).Run(); err != nil {
		log.Fatalf("error executing gofmt: %s", err)
	}

	// to catch everything else (hopefully)
	if err := exec.Command("go", "vet", "./...").Run(); err != nil {
		log.Fatalf("error executing go vet: %s", err)
	}
}
