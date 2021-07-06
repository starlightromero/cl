package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

var (
	c = colly.NewCollector()
)

func main() {
	args := os.Args

	if len(args) == 1 {
		displayHelp()
		return
	}

	switch args[1] {
	case "area":
		getAreas()
	case "search":
		searchCraigslist()
	}

}

func displayHelp() {
	fmt.Println("A all-purpose search engine for Craigslist free items")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("  area       Get a list of all Craigslist areas")
	fmt.Println("  search     Search a Craigslist area for free items")
	fmt.Println("")
	fmt.Println("Run 'cl COMMAND --help' for more information on a command.")
}

func getAreas() {

	c.OnHTML("h4+ul>li>a", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	link := "https://www.craigslist.org/about/sites"

	c.Visit(link)
}

func searchCraigslist() {
	var area string
	var query string

	flag.StringVar(&area, "a", "", "craigslist area to search")
	flag.StringVar(&area, "area", "", "craigslist area to search")
	flag.StringVar(&query, "q", "", "query to search")
	flag.StringVar(&query, "query", "", "query to search")

	flag.Parse()

	getItems(area, query)
}

func getItems(area, query string) {
	c.OnHTML("h3.result-heading", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
		resultLink := e.ChildAttr("a", "href")
		getItemDescription(resultLink)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	link := fmt.Sprintf("https://%s.craigslist.org/d/free-stuff/search/zip?query=%s", area, query)

	c.Visit(link)
}

func getItemDescription(link string) {

	fmt.Println(link)

	// c.OnHTML("#postingbody", func(e *colly.HTMLElement) {
	// 	fmt.Println(e.Text)
	// 	return
	// })

	// c.Visit(link)
}
