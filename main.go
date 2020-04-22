package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

type Site struct {
	Url       string `json:"url"`
	Status    int    `json:"status_code"`
	Duration  int64  `json:"duration_ms"`
	Timestamp string `json:"timestamp"`
}

// OnHTML ..
func OnHTML(tag string, attr string, c *colly.Collector, q *queue.Queue) {
	c.OnHTML(fmt.Sprintf("%s[%s]", tag, attr), func(e *colly.HTMLElement) {
		link := e.Attr(attr)
		q.AddURL(e.Request.AbsoluteURL(link))
	})
}

func SiteLog(site Site, json_out bool) {
	if json_out {
		js, _ := json.Marshal(site)
		fmt.Println(string(js))
	} else {
		fmt.Printf("[%d] %5d ms : %s \n", site.Status, site.Duration, site.Url)
	}
}

// Visit ..
func Visit(url string, threads int, verbose bool, json_out bool) {
	c := colly.NewCollector()
	q, _ := queue.New(threads, &queue.InMemoryQueueStorage{MaxSize: 1000})

	mainSite := Site{}

	OnHTML("link", "href", c, q)
	OnHTML("script", "src", c, q)
	OnHTML("img", "src", c, q)

	c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put(r.URL.String(), time.Now())
	})

	c.OnResponse(func(r *colly.Response) {
		u := r.Request.URL.String()
		start := r.Ctx.GetAny(u)
		if start != nil {
			duration := time.Now().Sub(start.(time.Time))
			site := Site{u, r.StatusCode, duration.Milliseconds(), time.Now().Format(time.RFC3339)}

			if u == url {
				mainSite = site
			}

			if verbose {
				SiteLog(site, json_out)
			}
		}
	})

	q.AddURL(url)
	start := time.Now()
	q.Run(c)
	duration := time.Now().Sub(start)
	mainSite.Duration = duration.Milliseconds()

	SiteLog(mainSite, json_out)
}

type Args struct {
	Threads int
	Verbose bool
	Json    bool
	Urls    []string
}

func usage() (Args, bool) {
	args := Args{}
	args.Urls = []string{}

	flag.IntVar(&args.Threads, "t", 1, "thread pool count")
	flag.BoolVar(&args.Verbose, "v", false, "verbose")
	flag.BoolVar(&args.Json, "json", false, "json output")
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
		Visit(url, args.Threads, args.Verbose, args.Json)

		if args.Verbose {
			fmt.Println()
		}
	}
}
