package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://www.marmiton.org/recettes/recette_tiramisu-au-carambar_56502.aspx", nil)

	check(err)

	req.Header.Set("User-Agent", "https://github.com/magleff/learn-to-boil")
	res, err2 := client.Do(req)

	check(err2)

	defer res.Body.Close()
	bytes, _ := ioutil.ReadAll(res.Body)

	fmt.Println("HTML:\n\n", string(bytes))
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
