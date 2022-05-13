package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"strconv"
)

//will traverse all of the dead links on the dsu cs cit page and print them to the screen

func citScrapper() {
	depth, err := getInput("How deep would you like to search?")
	if err != nil {
		// fmt.Println(err)
		log.Fatal(err)
	}
	intDepth, err := strconv.Atoi(depth)
	if err != nil {
		log.Fatal(err)
	}
	// finMap := make(map[string]bool)
	// dead := make([]string,1)
	// parentMap := make(map[string]string)
	dead := make(map[string]error)
	c := colly.NewCollector(
		colly.AllowedDomains("cit.dixie.edu", "https://cit.dixie.edu/", "www.cit.dixie.edu"),
		colly.MaxDepth(intDepth),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		link := e.Attr("href")
		// if !strings.HasPrefix(link, "/title") {
		// 	return
		// }
		// parentMap[link] = e.Request.URL.String()
		e.Request.Visit(link)
	})

	c.OnRequest(func(r *colly.Request) {
		// finMap[r.URL.String()] = true
		log.Println("visiting", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		// finMap[r.Request.URL.String()] = false
	})

	c.OnError(func(r *colly.Response, err error) {
		// fmt.Println("something went wrong with ",r.Request.URL.String())
		// dead = append(dead, r.Request.URL.String())
		fmt.Println(err)
		dead[r.Request.URL.String()] = err
	})

	c.Visit("https://cit.dixie.edu/cs/")

	fmt.Printf("\nDONE\n\n")

	for key, val := range dead {
		fmt.Printf("The link %v is dead. The error was %v.\n", key, val)
	}

	fmt.Println(c)
}
