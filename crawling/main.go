package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

// Selects every target nodes from selector node in document d
func Select(d *goquery.Document, selector string, target string) []string {
	infos := []string{}
	d.Find(selector).Each(func(i int, s *goquery.Selection) {
		t := s.Find(target).Text()
		infos = append(infos, t)
	})
	return infos
}

// Returns document from url string
func Scrape(url string) *goquery.Document {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}
	return doc
}

func main() {
	containerSelector := "#charts > div > div.chart-list.container > ol > li"
	songSelector := "button > span.chart-element__information > span.chart-element__information__song.text--truncate.color--primary"
	artistSelector := "button > span.chart-element__information > span.chart-element__information__artist.text--truncate.color--secondary"
	rankSelector := "button > span.chart-element__rank.flex--column.flex--xy-center.flex--no-shrink > span.chart-element__rank__number"

	doc := Scrape("https://www.billboard.com/charts/hot-100")
	songs := Select(doc, containerSelector, songSelector)
	artists := Select(doc, containerSelector, artistSelector)
	ranks := Select(doc, containerSelector, rankSelector)

	for i := 0; i < len(songs); i++ {
		fmt.Printf("%s: %s - %s\n", ranks[i], artists[i], songs[i])
	}
}
