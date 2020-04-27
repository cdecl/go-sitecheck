package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

type Site struct {
	Url        string `json:"url"`
	Status     int    `json:"status_code"`
	Duration   int64  `json:"duration_ms"`
	ContentLen uint64 `json:"content_length"`
	Timestamp  string `json:"timestamp"`
}

func siteLogging(site Site, jsonout bool) {
	if jsonout {
		js, _ := json.Marshal(site)
		fmt.Println(string(js))
	} else {
		fmt.Printf("[%3d] %7s [%4d ms] %s \n", site.Status, humanize.Bytes(site.ContentLen), site.Duration, site.Url)
	}
}

// OnHTML ..
func OnHTML(tag string, attr string, c *colly.Collector, q *queue.Queue) {
	tagInfo := fmt.Sprintf("%s[%s]", tag, attr)

	c.OnHTML(tagInfo, func(e *colly.HTMLElement) {
		link := e.Attr(attr)
		rel := e.Attr("rel")

		if tag == "link" && rel == "stylesheet" {
			q.AddURL(e.Request.AbsoluteURL(link))
		} else if tag != "link" {
			q.AddURL(e.Request.AbsoluteURL(link))
		}
	})
}

// OnRequest ..
func OnRequest(c *colly.Collector) {
	c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put(r.URL.String(), time.Now())
	})
}

// OnRequest ..
func OnResponse(mainSite *Site, totcontentlength *uint64, args Args, c *colly.Collector, q *queue.Queue) {
	c.OnResponse(func(r *colly.Response) {
		u := r.Request.URL.String()
		start := r.Ctx.GetAny(u)
		contentlength := uint64(len(r.Body))
		*totcontentlength += contentlength

		if start != nil {
			duration := time.Now().Sub(start.(time.Time))
			site := Site{u, r.StatusCode, duration.Milliseconds(), contentlength, time.Now().Format(time.RFC3339)}

			if u == mainSite.Url {
				*mainSite = site
			}

			if args.Verbose {
				siteLogging(site, args.Json)
			}
		}
	})
}

// OnError ..
func OnError(mainSite *Site, totcontentlength *uint64, args Args, c *colly.Collector) {
	c.OnError(func(r *colly.Response, err error) {
		u := r.Request.URL.String()
		start := r.Ctx.GetAny(u)
		contentlength := uint64(len(r.Body))
		*totcontentlength += contentlength

		duration := time.Now().Sub(start.(time.Time))
		site := Site{u, r.StatusCode, duration.Milliseconds(), contentlength, time.Now().Format(time.RFC3339)}

		if u == mainSite.Url {
			*mainSite = site
		}

		if args.Verbose {
			siteLogging(site, args.Json)
		}
	})
}

// Visit ..
func Visit(url string, args Args) {
	c := colly.NewCollector(
		colly.MaxDepth(1),
	)
	c.SetRequestTimeout(time.Duration(args.Timeout) * time.Second)

	q, _ := queue.New(args.Threads, &queue.InMemoryQueueStorage{MaxSize: 1000})

	mainSite := Site{}
	mainSite.Url = url
	totcontentlength := uint64(0)

	OnHTML("link", "href", c, q)
	OnHTML("script", "src", c, q)
	OnHTML("img", "src", c, q)

	OnRequest(c)
	OnResponse(&mainSite, &totcontentlength, args, c, q)
	OnError(&mainSite, &totcontentlength, args, c)

	q.AddURL(url)
	start := time.Now()
	q.Run(c)
	duration := time.Now().Sub(start)
	mainSite.Duration = duration.Milliseconds()
	mainSite.ContentLen = totcontentlength

	siteLogging(mainSite, args.Json)
}

type Args struct {
	Threads       int
	Verbose       bool
	Json          bool
	ContentLength int64
	Timeout       int
	Urls          []string
}

func usage() (Args, bool) {
	args := Args{}
	args.Urls = []string{}

	flag.IntVar(&args.Threads, "t", 1, "thread pool count")
	flag.BoolVar(&args.Verbose, "v", false, "verbose")
	flag.BoolVar(&args.Json, "json", false, "json output")
	flag.IntVar(&args.Timeout, "m", 10, "request timeout sec")
	flag.Bool("", false, "ver. 200427")
	flag.Parse()

	args.Urls = append(args.Urls, flag.Args()...)

	if len(args.Urls) == 0 {
		flag.Usage()
		return args, false
	}
	return args, true
}

func main() {
	args, ok := usage()
	if !ok {
		return
	}

	for _, url := range args.Urls {
		Visit(url, args)

		if args.Verbose {
			fmt.Println()
		}
	}
}
