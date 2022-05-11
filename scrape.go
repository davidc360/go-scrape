package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	var tickers []string
	var companyNames []string

	c.OnHTML(".simpTblRow > td > a", func(e *colly.HTMLElement) {
		// fmt.Printf("Most popular ticker: %s\n", e.Text)
		tickers = append(tickers, e.Text)
	})

	c.OnHTML("[aria-label=\"Name\"]", func(e *colly.HTMLElement) {
		// fmt.Printf("Most popular company: %s\n", e.Text)
		companyNames = append(companyNames, e.Text)
	})

	c.Visit("https://finance.yahoo.com/trending-tickers")

	fmt.Printf("The most popular ticker is $%s, %s\n", tickers[0], companyNames[0])
}