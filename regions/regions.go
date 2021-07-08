package regions

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/starlightromero/cl/help"
)

func Get(wantHelp bool, c *colly.Collector) {
	help.Check(wantHelp, printHelp)

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	link := "https://www.craigslist.org/about/sites"

	c.Visit(link)

}

func Check(region string) int {
	var n int

	switch region {
	case "US":
		n = 1
	case "Canada":
		n = 2
	case "Europe":
		n = 3
	case "Asia, Pacific and Middle East":
		n = 4
	case "Oceania":
		n = 5
	case "Latin America and Caribbean":
		n = 6
	case "Africa":
		n = 7
	}

	return n
}

func printHelp() {
	fmt.Println("")
	fmt.Println("Usage: cl regions")
	fmt.Println("")
	fmt.Println("Get Craigslist regions")
}
