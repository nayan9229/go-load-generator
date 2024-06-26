package main

import (
	"context"
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"
	"time"
	"visiter/load"

	"math/rand"

	"github.com/chromedp/chromedp"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: helloserver [options]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

var (
	greeting = flag.String("g", "Hello", "Greet with `greeting`")
	addr     = flag.String("addr", "localhost:8080", "address to serve")
)

func main() {
	// Parse flags.
	flag.Usage = usage
	flag.Parse()

	// Parse and validate arguments (none).
	args := flag.Args()
	if len(args) != 0 {
		usage()
	}

	rand.Seed(time.Now().UnixNano())

	// Register handlers.
	// All requests not otherwise mapped with go to greet.
	// /version is mapped specifically to version.
	http.HandleFunc("/", greet)
	http.HandleFunc("/version", version)
	http.HandleFunc("/visit", visit)
	http.HandleFunc("/load", loadTest)

	log.Printf("serving http://%s\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func loadTest(w http.ResponseWriter, r *http.Request) {
	// Get query parameters
	timeout, err := strconv.Atoi(r.URL.Query().Get("timeout"))
	if err != nil {
		http.Error(w, "Invalid timeout parameter", 400)
		return
	}

	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "Missing url parameter", 400)
		return
	}

	runtime, err := strconv.Atoi(r.URL.Query().Get("runtime"))
	if err != nil {
		http.Error(w, "Invalid runtime parameter", 400)
		return
	}

	parallelRequests, err := strconv.Atoi(r.URL.Query().Get("parallel_requests"))
	if err != nil {
		http.Error(w, "Invalid parallel_requests parameter", 400)
		return
	}

	// Start load test
	response := load.StartLoadTest(url, int(timeout), int(runtime), parallelRequests)
	// Write response
	fmt.Fprintf(w, response)
}

func version(w http.ResponseWriter, r *http.Request) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		http.Error(w, "no build information available", 500)
		return
	}

	fmt.Fprintf(w, "<!DOCTYPE html>\n<pre>\n")
	fmt.Fprintf(w, "%s\n", html.EscapeString(info.String()))
}

func greet(w http.ResponseWriter, r *http.Request) {
	name := strings.Trim(r.URL.Path, "/")
	if name == "" {
		name = "Gopher"
	}

	fmt.Fprintf(w, "<!DOCTYPE html>\n")
	fmt.Fprintf(w, "%s, %s!\n", *greeting, html.EscapeString(name))
}

func visit(w http.ResponseWriter, r *http.Request) {
	url := "https://demo.begenuin.com/demo/qa-demo.html?tag_id=2874"
	fmt.Println(url)
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 5)

	for i := 0; i < 20; i++ {
		time.Sleep(4 * time.Second)
		wg.Add(1)
		semaphore <- struct{}{}

		go func(i int) {
			defer wg.Done()
			defer func() { <-semaphore }()
			fmt.Printf("Starting for %d \n", i)
			err := OpenAndInteractWithPage(url, i)
			fmt.Printf("Completed for %d \n", i)
			if err != nil {
				log.Printf("Error in goroutine %d: %v", i, err)
			}
		}(i)
	}

	wg.Wait()
	fmt.Fprintf(w, "All tasks completed")
}

func OpenAndInteractWithPage(url string, ID int) error {
	// Create a new context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Create a timeout context
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	var buf []byte
	sleep := randRange(5, 45)
	fmt.Printf("sleep: %v\n", sleep)
	// Run tasks
	err := chromedp.Run(ctx,
		// Navigate to the URL
		chromedp.Navigate(url),
		// Wait for the network to be idle
		chromedp.WaitReady("body", chromedp.ByQuery),

		chromedp.Sleep(time.Duration(sleep)*time.Second),
		// Take a screenshot
		chromedp.FullScreenshot(&buf, 90),
		// // Scroll to bottom
		// chromedp.Evaluate(`window.scrollTo(0, document.body.scrollHeight)`, nil),
		// // Wait for a short duration to ensure the scroll completes
		// chromedp.Sleep(2*time.Second),
		// // Scroll back to top
		// chromedp.Evaluate(`window.scrollTo(0, 0)`, nil),
		// Wait again to ensure the scroll completes
		// chromedp.Sleep(2*time.Second),
	)

	if err != nil {
		log.Fatal(err)
	}

	// Save the screenshot to a file
	if err := saveScreenshot(buf, ID); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Screenshot taken and page scrolled successfully")
	return nil
}

func saveScreenshot(buf []byte, ID int) error {
	return os.WriteFile(fmt.Sprintf("screenshot_%d.png", ID), buf, 0644)
}

func randRange(min, max int) int {
	return rand.Intn(max+1-min) + min
}
