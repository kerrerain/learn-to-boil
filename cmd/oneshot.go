package cmd

import (
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
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

func init() {
	RootCmd.AddCommand(oneshotCmd)
}

var oneshotCmd = &cobra.Command{
	Use:   "oneshot",
	Short: "...",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		var lines []string

		location, err := randomLocation()

		check(err)

		log.Println("Analyzing random recipe:", location)

		response, _ := request(location)
		root, err := html.Parse(response.Body)

		check(err)

		ingredients, ok := scrape.Find(root, scrape.ByClass("m_content_recette_ingredients"))
		if ok {
			lines = ExtractLines(scrape.Text(ingredients))
			log.Println("Extracting ingredients", lines)
		}

		defer response.Body.Close()

		// Open the my.db data file in your current directory.
		// It will be created if it doesn't exist.
		db, err := bolt.Open("data.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
		if err != nil {
			log.Fatal(err)
		}

		check(err)

		defer db.Close()

		db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte("Ingredients"))
			if err != nil {
				return fmt.Errorf("create bucket: %s", err)
			}
			return nil
		})

		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("Ingredients"))
			err := b.Put([]byte(location), []byte(strings.Join(lines, "#")))
			return err
		})

		db.View(func(tx *bolt.Tx) error {
			// Assume bucket exists and has keys
			b := tx.Bucket([]byte("Ingredients"))

			b.ForEach(func(k, v []byte) error {
				fmt.Printf("key=%s, value=%s\n", k, v)
				return nil
			})
			return nil
		})
	},
}
