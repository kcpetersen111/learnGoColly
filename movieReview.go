package main

import (
	"bufio"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
	"strings"
	// "encoding/json"
)

var first bool

//will recive a prompt to print to the screen and then get the responce from the user
func getInput(prompt string) (string, error) {
	fmt.Printf("%s\n", prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		return "", fmt.Errorf("Error %v while reading from input", err)
	}
	return scanner.Text(), nil
}

type Movie struct {
	title  string
	rating string
}

//finds the rating of a movie given by user input
func movieReview() {
	movie, err := getInput("What movie review would you like to find?")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	c := colly.NewCollector(
		// colly.AllowedDomains("imdb.com"),
		colly.MaxDepth(2),
		// colly.AllowURLRevisit(false),
	)

	finalMap := make(map[string]Movie)

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		link := e.Attr("href")

		if !strings.HasPrefix(link, "/title") {
			return
		}

		e.Request.Visit(link)
	})

	//gets the name of the movies
	c.OnHTML("h1", func(e *colly.HTMLElement) {
		currURL := e.Request.URL.String()
		if !strings.Contains(currURL, "/title") {
			return
		}
		// header := e.Attr("h1")
		// fmt.Printf(" Movie: %v\n",e.Text)
		temp := Movie{
			title:  e.Text,
			rating: "",
		}
		finalMap[currURL] = temp
	})

	//gets the rating of the movie
	c.OnHTML("div[class='sc-7ab21ed2-2 kYEdvH']", func(e *colly.HTMLElement) {
		currURL := e.Request.URL.String()
		if !strings.Contains(currURL, "/title") {
			return
		}
		// header := e.Attr("h1")
		// fmt.Printf(" Movie: %v\n",e.Text)
		temp := Movie{
			title:  finalMap[currURL].title,
			rating: e.Text,
		}
		finalMap[currURL] = temp
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})

	mvurl := "https://www.imdb.com/find?q="
	tmp := strings.Split(movie, " ")
	tmp2 := strings.Join(tmp, "+")
	mvurl += tmp2
	fmt.Println(mvurl)
	c.Visit(mvurl)
	// fmt.Println(mvurl)

	fmt.Println(c)

	// fmt.Println(finalMap)
	for _, val := range finalMap {
		fmt.Println(val)
	}

}
