package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gocolly/colly"
)

//will traverse all of the dead links on the dsu cs cit page and print them to the screen

func fasterCitScrapper() {
	temp := bufio.NewScanner(os.Stdin)
	fmt.Println("How deep would you like to search?")
	succ := temp.Scan()
	if !succ {
		fmt.Println("no input supplied")
		return
	}
	depth := temp.Text()
	counter := 0
	// depth, err := getInput("How deep would you like to search?")
	// if err != nil {
	// 	// fmt.Println(err)
	// 	log.Fatal(err)
	// }
	intDepth, err := strconv.Atoi(depth)
	if err != nil {
		log.Fatal(err)
	}
	// finMap := make(map[string]bool)
	// dead := make([]string,1)
	// parentMap := make(map[string]string)
	dead := make(map[string]error)
	c := colly.NewCollector(
		// colly.AllowedDomains("computing.utahtech.edu", "https://computing.utahtech.edu/", "www.cit.dixie.edu", ".edu"),
		colly.MaxDepth(intDepth),
		colly.Async(true),
	)
	c.Limit(&colly.LimitRule{Parallelism: 10})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		counter++

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

	c.Visit("https://computing.utahtech.edu/")

	c.Wait()
	fmt.Printf("\nDONE\n\n")

	for key, val := range dead {
		fmt.Printf("The link %v is dead. The error was %v.\n", key, val)
	}
	fmt.Println(c)
	fmt.Printf("There was a grand total of %d number of links\n", counter)
}
