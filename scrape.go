package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	c.OnHTML(".simpTblRow > td > a", func(e *colly.HTMLElement) {
		fmt.Printf("Most popular ticker: %s\n", e.Text)
	})

	c.Visit("https://finance.yahoo.com/trending-tickers")
}