package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const BASE_URL = "https://docs.github.com"

const OUTPUT_FILE_HEADER = `
package mock


`
const OUTPUT_FILEPATH = "src/mock/endpointpattern.go"

type ScrapeResult struct {
	HttpMethod      string
	EndpointPattern string
}

func getGoQueryDocumentFromUrl(urlToScrape string) (*goquery.Document, error) {
	res, err := http.Get(urlToScrape)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return doc, err
}

func scrapeReferenceTab() []string {
	doc, err := getGoQueryDocumentFromUrl("https://docs.github.com/en/rest")

	if err != nil {
		log.Fatal(err)
	}

	restEndpoints := []string{}

	doc.Find("#__next > div > div > nav > ul > li:nth-child(3) > ul > li:nth-child(2) > details > ul > li > a").
		Each(func(i int, s *goquery.Selection) {
			if v, ok := s.Attr("href"); ok {
				restEndpoints = append(restEndpoints, fmt.Sprintf("%s%s", BASE_URL, v))
			}
		})

	return restEndpoints
}

func scrapeApiReference(url string) <-chan ScrapeResult {
	scrapedEndpoints := make(chan ScrapeResult)

	go func() {
		doc, err := getGoQueryDocumentFromUrl(url)
		if err != nil {
			log.Fatal(err)
		}

		doc.Find("#article-contents > div > pre > code").Each(
			func(i int, s *goquery.Selection) {
				combinedInnerTextSplit := strings.Split(
					s.Text(),
					" ",
				)

				scrapedEndpoints <- ScrapeResult{
					HttpMethod:      combinedInnerTextSplit[0],
					EndpointPattern: combinedInnerTextSplit[1],
				}
			},
		)

		close(scrapedEndpoints)
	}()

	return scrapedEndpoints
}

func formatToGolangVarName(sr ScrapeResult) string {
	epSplit := strings.Split(strings.ReplaceAll(sr.EndpointPattern, "-", "/"), "/")

	result := ""

	for _, part := range epSplit {
		if len(part) < 1 || string(part[0]) == "{" {
			continue
		}

		result = result + strings.Title(part)
	}

	return strings.Title(sr.HttpMethod) + result
}

func formatToGolangVarNameAndValue(sr ScrapeResult) string {
	return fmt.Sprintf(
		"var %s EndpointPattern = \"%s\"\n",
		formatToGolangVarName(sr),
		sr.EndpointPattern,
	)
}

func main() {
	referenceUrlsToScrape := scrapeReferenceTab()

	buf := bytes.NewBuffer([]byte(OUTPUT_FILE_HEADER))

	for _, url := range referenceUrlsToScrape {
		res := scrapeApiReference(url)

		for r := range res {
			buf.WriteString(formatToGolangVarNameAndValue(r))
		}
	}

	ioutil.WriteFile(
		OUTPUT_FILEPATH,
		buf.Bytes(),
		0755,
	)
}
