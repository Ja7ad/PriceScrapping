package src

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"github.com/gocolly/colly"
)

const xpath = "xpath.json"

type Prices struct {
	Dollar string `json:"dollar"`
	Euro   string `json:"euro"`
	Coin   string `json:"coin"`
	Gold   string `json:"gold"`
}

func Scrapper(prc *Prices) map[string]string {
	// Load XPATH
	XpathFile, err := os.Open(xpath)
	if err != nil {
		panic(err)
	}

	var price Prices
	byteValue, _ := ioutil.ReadAll(XpathFile)
	if err := json.Unmarshal(byteValue, &price); err != nil {
		panic(err)
	}

	// Create Collector Instance
	collect := colly.NewCollector()

	// Collect Data
	collect.OnHTML(price.Dollar, func(e *colly.HTMLElement) {
		prc.Dollar = e.Text
	})
	collect.OnHTML(price.Euro, func(e *colly.HTMLElement) {
		prc.Euro = e.Text
	})
	collect.OnHTML(price.Gold, func(e *colly.HTMLElement) {
		prc.Gold = e.Text
	})
	collect.OnHTML(price.Coin, func(e *colly.HTMLElement) {
		prc.Coin = e.Text
	})

	// Visit Target
	_ = collect.Visit("https://www.tgju.org")

	return map[string]string{
		"dollar" : prc.Dollar,
		"euro" : prc.Euro,
		"gold" : prc.Gold,
		"coin" : prc.Coin,
	}
}