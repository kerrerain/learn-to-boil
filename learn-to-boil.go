package main

import (
	"fmt"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://www.marmiton.org/recettes/recette_roule-au-saumon-et-aux-epinards_30443.aspx", nil)

	check(err)

	// Read http://blog.mischel.com/2011/12/20/writing-a-web-crawler-politeness/
	req.Header.Set("User-Agent", "https://github.com/magleff/learn-to-boil")
	res, err := client.Do(req)

	check(err)

	defer res.Body.Close()

	root, err := html.Parse(res.Body)

	check(err)

	// Search for the title
	ingredients, ok := scrape.Find(root, scrape.ByClass("m_content_recette_ingredients"))
	if ok {
		// Print the title
		fmt.Println(scrape.TextJoin(ingredients, func(chunks []string) string {
			for _, chunk := range chunks {
				fmt.Println(chunk)
			}
			return strings.Join(chunks, " ")
		}))
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ExtractLines(input string) []string {
	r, _ := regexp.Compile(`<[^>]*>`)
	str := r.ReplaceAllString(input, "")

	extract, _ := regexp.Compile(`(-.+\n?[\s]+[^-]+)`)
	lines := extract.FindAllString(str, -1)

	eraseWhitespaces, _ := regexp.Compile(`\s{2,}`)

	for index, _ := range lines {
		lines[index] = eraseWhitespaces.ReplaceAllString(lines[index], " ")
		lines[index] = strings.Replace(lines[index], "- ", "", -1)
		lines[index] = strings.TrimSpace(lines[index])
	}

	return lines
}

func ExtractIngredients(input string) []string {
	return []string{}
}
