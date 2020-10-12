package main

import (
	"github.com/gocolly/colly"
	"strings"
)

// Export Scraper function
func Scrape(tt string) []string{

	var quotes []string

	c := colly.NewCollector()

	c.OnHTML("div.sodatext", func(sodaSection *colly.HTMLElement){
		sodaSection.ForEach("p", func(_ int, quote *colly.HTMLElement) {

			// Trim WhiteSpaces and remove \n chars
			quotes = append(quotes, strings.TrimSpace(strings.ReplaceAll(quote.Text, "\n", "")))

		})
		quotes = append(quotes, "<br>")
	})

	URL := "https://www.imdb.com/title/" + tt + "/quotes"
	c.Visit(URL)

	return quotes
}