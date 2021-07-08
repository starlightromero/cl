package search

import (
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
	"github.com/starlightromero/cl/help"
)

func Execute(area, query string, today bool, wantHelp bool, c *colly.Collector) {
	help.Check(wantHelp, printHelp)

	var postedToday int

	if area == "" && query == "" {
		printErrorHelp()
		os.Exit(1)
	}

	if today {
		postedToday = 1
	}

	c.OnHTML("h3.result-heading", func(e *colly.HTMLElement) {
		title := strings.TrimSpace(e.Text)
		link := e.ChildAttr("a", "href")

		fmt.Println(title)
		fmt.Println(link)
		fmt.Println("")
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	link := fmt.Sprintf("https://%s.craigslist.org/d/free-stuff/search/zip?query=%s&postedToday=%d", area, query, postedToday)

	c.Visit(link)
}

func printHelp() {
	fmt.Println("")
	fmt.Println("Usage: cl search [OPTIONS]")
	fmt.Println("")
	fmt.Println("Search Craigslist for free items")
	fmt.Println("Options:")
	fmt.Println("  -a, --area string   The area to search for items within")
	fmt.Println("  -q, --query string  The query to search")
	fmt.Println("  -t, --today bool    Only show items that were posted today")
}

func printErrorHelp() {
	fmt.Println("\"cl search\" requires at least 1 flag.")
	fmt.Println("See 'cl search --help'.")
	fmt.Println("")
	fmt.Println("Usage: cl search [OPTIONS]")
	fmt.Println("")
	fmt.Println("Search Craigslist for free items")
}
