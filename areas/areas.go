package areas

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
	"github.com/starlightromero/cl/help"
	"github.com/starlightromero/cl/regions"
)

func Get(region string, wantHelp bool, c *colly.Collector) {
	help.Check(wantHelp, printHelp)

	n := regions.Check(region)

	if region != "" {
		selector := fmt.Sprintf("h1:nth-of-type(%d)+div h4+ul>li>a", n)
		c.OnHTML(selector, func(e *colly.HTMLElement) {
			href := stripLink(e.Attr("href"))
			fmt.Println(href)
		})
	} else {
		c.OnHTML("h4+ul>li>a", func(e *colly.HTMLElement) {
			href := stripLink(e.Attr("href"))
			fmt.Println(href)
		})
	}

	link := "https://www.craigslist.org/about/sites"

	c.Visit(link)
}

func stripLink(link string) string {
	link = strings.TrimPrefix(link, "https://")
	link = strings.Split(link, ".")[0]
	return link
}

func printHelp() {
	fmt.Println("")
	fmt.Println("Usage: cl areas [OPTIONS]")
	fmt.Println("")
	fmt.Println("Get Craigslist areas")
	fmt.Println("")
	fmt.Println("Options:")
	fmt.Println("  -r, --region string      The region to get all areas within")
}
