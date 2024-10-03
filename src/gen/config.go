package gen

const GITHUB_OPENAPI_DEFINITION_LOCATION = "https://raw.githubusercontent.com/github/rest-api-description/refs/heads/main/descriptions/api.github.com/api.github.com.json"
const GITHUB_OPENAPI_ENTERPRISE_DEFINITION_LOCATION = "https://raw.githubusercontent.com/github/rest-api-description/main/descriptions/ghec/ghec.json"

const OUTPUT_FILE_HEADER = `package mock

// Code generated; DO NOT EDIT.

`
const OUTPUT_DIR = "src/mock"

type ConfigEntry struct {
	URL    string
	Prefix string
}

const BASE_FILENAME = "endpointpattern.go"

// Computes the filename based on the prefix
func (c ConfigEntry) ComputeFilename() string {
	return c.Prefix + BASE_FILENAME
}

func (c ConfigEntry) FormatPrefix() string {
	return Caser.String(c.Prefix)
}
