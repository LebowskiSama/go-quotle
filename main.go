package main

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

// Declare MiddleWare to use and concurrently set CORS headers
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// Export Scraper function
func Scrape(tt string) []string {

	var quotes []string

	c := colly.NewCollector()

	c.OnHTML("div.sodatext", func(sodaSection *colly.HTMLElement) {
		sodaSection.ForEach("p", func(_ int, quote *colly.HTMLElement) {

			// Trim WhiteSpaces and remove \n chars
			quotes = append(quotes, strings.TrimSpace(strings.ReplaceAll(quote.Text, "\n", "")))

		})
		quotes = append(quotes, "<hr>")
	})

	URL := "https://www.imdb.com/title/" + tt + "/quotes"
	c.Visit(URL)

	return quotes
}

func main() {

	// Run a gin server
	r := gin.Default()

	// Use middleware function
	r.Use(CORSMiddleware())

	// Endpoint to serve JSON
	r.GET("/movie/", serve)

	// Initialize gin
	r.Run()
}

func serve(c *gin.Context) {

	// Get the ImdbID
	tt := c.Query("tt")

	quotes := Scrape(tt)
	c.JSON(200, quotes)
}
