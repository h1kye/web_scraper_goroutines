package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"strconv"
	"strings"
	"time"
)

// -----------------------
// WEBSCRAPER - GOROUTINES [BASIC CHANNELS -> NO WAITGROUPS!!]
// -----------------------
// GOROUTINE - EXTRACT MULTIPLE URLS, FROM LIST THEN RUN IN PARALLEL
// SCRAPER MOD - https://github.com/gocolly/colly
// GOROUTINE SCRAPER EXAMPLE - https://medium.com/@emiliocliff/unleashing-concurrency-in-web-scraping-with-go-routines-a-beginners-guide-371adbf912a5

func main() {

	// URL LIST
	urlList := []string{"wirefast.com", "itv.com", "sky.com", "channel4.com"}

	// CHANNEL CREATE
	resultsChan := make(chan string)

	//TIME EXECUTION
	startTime := time.Now()

	// SCRAPER RUN
	for _, siteName := range urlList {
		// > JOIN "HTTPS://" TO SITE URL
		fullURL := strings.Join([]string{"https://", siteName}, "")

		// GOROUTINE SCRAPER RUN - SEND RESULTS TO "CHANNEL"
		go scraperUrl(fullURL, resultsChan)
	}

	// CHANNEL RECEIVE - PRINT RESULTS
	fmt.Printf(<-resultsChan)

	// TOTAL DURATION RUN
	durationRun := time.Since(startTime)
	fmt.Printf("\nTOTAL RUNTIME - %v ns - %.2f seconds\n", durationRun, durationRun.Seconds())
}

// URL SCRAPE
// -----------
//
//	 [ARGS]: STRING -> url
//			 CHANNEL -> "send" data to channel (outbound)
func scraperUrl(site string, chSexy chan<- string) {

	// CREATE HANDLER TO PULL URL!!
	c := colly.NewCollector()

	fmt.Printf("This is the beginning URL ..... %v\n", site)

	// OnResponse callback - RETURN PAGE CONTENTS
	c.OnResponse(func(r *colly.Response) {

		// DISPLAY - FULL PAGE CONTENTS
		pageFullContents := string(r.Body)
		fmt.Printf("FULL PAGE CONTENTS - %v\n %v\n", pageFullContents)

		// DISPLAY - "RETURN CODE"
		statusCode := r.StatusCode
		fmt.Printf("STATUS CODE - %v\n", statusCode)

		// SEND TO CHANNEL - RETURN CODE (CONVERT FROM INT TO STRING)
		chSexy <- strconv.Itoa(statusCode)
	})

	// CONNECT TO SITE
	//
	// > CONNECT TO SITE URL
	result := c.Visit(site)

	// CHECK CONNECTION STATUS
	if result != nil {
		fmt.Printf("Error extracting page - %s", result)
		//continue
		log.Fatal(result)
	}
}

// [MARK]
// VERSION 1 - GROUTINE (NO WAITGROUPS!!) -> BASIC SETUP
//
// to do --> intorduce "waitgroups" --> more complex channel management
