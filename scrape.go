package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"github.com/gocolly/colly"
)

type Data struct {
	Ticker string
	CompanyName string
}

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

	data := Data{Ticker: tickers[0], CompanyName: companyNames[0]}
	fmt.Println(data)

	jsonData, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile("most-popular-ticker.json", jsonData, 0644)
}