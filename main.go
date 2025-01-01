package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"strings"
	"time"
)

// WEBSCRAPER - GOROUTINES
// -----------------------
// GOROUTINE - EXTRACT MULTIPLE URLS, FROM LIST THEN RUN IN PARALLEL
// SCRAPER MOD - https://github.com/gocolly/colly
// GOROUTINE SCRAPER EXAMPLE - https://medium.com/@emiliocliff/unleashing-concurrency-in-web-scraping-with-go-routines-a-beginners-guide-371adbf912a5

func main() {

	//TIME EXECUTION
	startTime := time.Now()

	// SCRAPER RUN
	scraperUrl()

	durationRun := time.Since(startTime)
	fmt.Printf("TOTAL RUNTIME - %v ns - %.2f seconds\n", durationRun, durationRun.Seconds())
}

// MARK -> HOW TO CREAT GOROUTINES + CHANNELS???

func scraperUrl() {

	// CREATE HANDLER TO PULL URL!!
	urlList := []string{"wirefast.com", "itv.com", "sky.com", "channel4.com"}
	c := colly.NewCollector()

	// LIST OF SITES - SCRAPE
	for _, site := range urlList {
		fmt.Printf("This is the beginning URL ..... %v\n", site)

		// OnResponse callback - RETURN PAGE CONTENTS
		c.OnResponse(func(r *colly.Response) {

			// DISPLAY - FULL PAGE CONTENTS
			pageFullContents := string(r.Body)
			fmt.Printf("FULL PAGE CONTENTS - %v \n", pageFullContents)

			// DISPLAY - "RETURN CODE"
			statusCode := r.StatusCode
			fmt.Printf("STATUS CODE - %v\n", statusCode)
		})

		// CONNECT TO SITE
		//
		// > JOIN "HTTPS://" TO SITE URL
		fullURL := strings.Join([]string{"https://", site}, "")
		// > CONNECT TO SITE URL
		result := c.Visit(fullURL)

		// CHECK CONNECTION STATUS
		if result != nil {
			fmt.Printf("Error extracting page - %s", result)
			//continue
			log.Fatal(result)
		}
	}
}
