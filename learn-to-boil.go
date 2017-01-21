package main

import (
	"errors"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func request(url string) (*http.Response, error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return errors.New("net/http: use last response")
		},
	}
	req, err := http.NewRequest("GET", url, nil)

	check(err)

	req.Header.Set("User-Agent", "https://github.com/magleff/learn-to-boil")

	return client.Do(req)
}

func randomLocation() (string, error) {
	response, err := request("http://www.marmiton.org/recettes/recette-hasard.aspx")

	if err != nil {
		location, _ := response.Location()
		return location.String(), nil
	} else {
		return "", errors.New("Random location not found.")
	}
}

func main() {
	location, err := randomLocation()

	check(err)

	log.Println("Analyzing random recipe:", location)

	response, _ := request(location)
	root, err := html.Parse(response.Body)

	check(err)

	ingredients, ok := scrape.Find(root, scrape.ByClass("m_content_recette_ingredients"))
	if ok {
		lines := ExtractLines(scrape.Text(ingredients))
		log.Println("Extracting ingredients", lines)
	}

	defer response.Body.Close()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ExtractLines(input string) []string {
	r, _ := regexp.Compile(`(-.[^-]+)`)
	lines := r.FindAllString(input, -1)

	for index, _ := range lines {
		lines[index] = strings.Replace(lines[index], "- ", "", -1)
		lines[index] = strings.TrimSpace(lines[index])
	}

	return lines
}
