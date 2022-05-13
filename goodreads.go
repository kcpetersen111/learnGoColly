package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"github.com/gocolly/colly"
	"strings"
)

func readPassword() (string, string) {
	// fmt.Println(ioutil.ReadFile("passwords.txt"))
	file, err := os.Open("passwords.txt")
	if err!= nil{
		log.Fatalf("Error opening file: %v\n",err)
	}
	scanner := bufio.NewScanner(file)
	temp := make([]string, 2)
	for scanner.Scan() {
		temp = append(temp, scanner.Text())
	}
	return temp[0], temp[1]
}

func goodreads() {
	username, password := readPassword()

	c := colly.NewCollector(
		colly.AllowURLRevisit(),
	)
	bookList := make([]string, 0)

	err := c.Post("https://www.goodreads.com/ap/signin?language=en_US&openid.assoc_handle=amzn_goodreads_web_na&openid.claimed_id=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0%2Fidentifier_select&openid.identity=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0%2Fidentifier_select&openid.mode=checkid_setup&openid.ns=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0&openid.pape.max_auth_age=0&openid.return_to=https%3A%2F%2Fwww.goodreads.com%2Fap-handler%2Fsign-in&siteState=63049c91a8055e8697b9a203dfa8cad1",
					map[string]string{"username":username, "password":password})
	if err != nil{
		log.Fatalf("Error loging into goodreads %v\n",err)
	}

	// attach callbacks after login
	c.OnResponse(func(r *colly.Response) {
		log.Println("response received", r.StatusCode)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})

	// c.OnHTML("A[href]", func(r *colly.HTMLElement){

	// 	link := r.Attr("href")
	// 	fmt.Println(link)
	// 	if !strings.HasPrefix(link, "/review") {
	// 		return
	// 	}

	// 	r.Request.Visit(link)
	// })

	c.OnHTML("A[href]", func(e *colly.HTMLElement){
		link := e.Attr("href")
		if !strings.HasPrefix(link, "/book/show")  {
			return
		}

		bookTitle := e.Attr("title")

		bookList = append(bookList, bookTitle)
		// fmt.Println(*(e.Response.Headers))

	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	// start scraping
	c.Visit("https://www.goodreads.com/review/list/35763417-kaleb-petersen?ref=nav_mybooks&shelf=to-read")

	for _,val := range bookList{
		fmt.Println(val)
	}

	fmt.Println(c)
}

