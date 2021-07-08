package subregions

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/starlightromero/cl/help"
	"github.com/starlightromero/cl/regions"
)

func Get(region string, wantHelp bool, c *colly.Collector) {
	help.Check(wantHelp, printHelp)

	n := regions.Check(region)

	if region != "" {
		selector := fmt.Sprintf("h1:nth-of-type(%d)+div h4", n)
		c.OnHTML(selector, func(e *colly.HTMLElement) {
			fmt.Println(e.Text)
		})
	} else {
		c.OnHTML("h4", func(e *colly.HTMLElement) {
			fmt.Println(e.Text)
		})
	}

	link := "https://www.craigslist.org/about/sites"

	c.Visit(link)

}

func printHelp() {
	fmt.Println("")
	fmt.Println("Usage: cl subregions [OPTIONS]")
	fmt.Println("")
	fmt.Println("Get Craigslist subregions")
	fmt.Println("")
	fmt.Println("Options:")
	fmt.Println("  -r, --region string  The region to get all areas within")
}
