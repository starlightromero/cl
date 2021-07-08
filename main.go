package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gocolly/colly"
	"github.com/starlightromero/cl/areas"
	"github.com/starlightromero/cl/regions"
	"github.com/starlightromero/cl/search"
	"github.com/starlightromero/cl/subregions"
)

var (
	c = colly.NewCollector()
)

func main() {
	var area string
	var region string
	var query string
	var today bool
	var help bool

	areasCmd := flag.NewFlagSet("areas", flag.ExitOnError)
	areasCmd.StringVar(&region, "r", "", "craigslist region to search")
	areasCmd.StringVar(&region, "region", "", "craigslist region to search")
	areasCmd.BoolVar(&help, "help", false, "get help")

	subregionsCmd := flag.NewFlagSet("subregions", flag.ExitOnError)
	subregionsCmd.StringVar(&region, "r", "", "craigslist region to search")
	subregionsCmd.StringVar(&region, "region", "", "craigslist region to search")
	subregionsCmd.BoolVar(&help, "help", false, "get help")

	regionsCmd := flag.NewFlagSet("regions", flag.ExitOnError)
	regionsCmd.BoolVar(&help, "help", false, "get help")

	searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
	searchCmd.StringVar(&area, "a", "", "craigslist area to search")
	searchCmd.StringVar(&area, "area", "", "craigslist area to search")
	searchCmd.StringVar(&query, "q", "", "query to search")
	searchCmd.StringVar(&query, "query", "", "query to search")
	searchCmd.BoolVar(&today, "t", false, "query for items posted today only")
	searchCmd.BoolVar(&today, "today", false, "query for items posted today only")
	searchCmd.BoolVar(&help, "help", false, "get help")

	helpCmd := flag.NewFlagSet("help", flag.ExitOnError)

	flag.Parse()

	if len(os.Args) < 2 {
		printHelp()
	}

	switch os.Args[1] {
	case "areas":
		areasCmd.Parse(os.Args[2:])
		areas.Get(region, help, c)
	case "subregions":
		subregionsCmd.Parse(os.Args[2:])
		subregions.Get(region, help, c)
	case "regions":
		regionsCmd.Parse(os.Args[2:])
		regions.Get(help, c)
	case "search":
		searchCmd.Parse(os.Args[2:])
		search.Execute(area, query, today, help, c)
	case "help":
		helpCmd.Parse(os.Args[2:])
		printHelp()
	default:
		printHelp()
	}

}

func printHelp() {
	fmt.Println("A all-purpose search engine for Craigslist free items")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("  areas        Get a list of all Craigslist areas")
	fmt.Println("  search       Search a Craigslist area for free items")
	fmt.Println("  subregions   Get a list of all Craigslist subregions")
	fmt.Println("  regions      Get a list of all Craigslist regions")
	fmt.Println("")
	fmt.Println("Run 'cl COMMAND --help' for more information on a command.")
	os.Exit(0)
}
