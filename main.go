package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

// OnHTML ..
func OnHTML(tag string, attr string, c *colly.Collector, q *queue.Queue) {
	c.OnHTML(fmt.Sprintf("%s[%s]", tag, attr), func(e *colly.HTMLElement) {
		link := e.Attr(attr)
		q.AddURL(e.Request.AbsoluteURL(link))
	})
}

// Visit ..
func Visit(url string, threads int, verbose bool) {
	c := colly.NewCollector()
	q, _ := queue.New(threads, &queue.InMemoryQueueStorage{MaxSize: 1000})

	OnHTML("link", "href", c, q)
	OnHTML("script", "src", c, q)
	OnHTML("img", "src", c, q)

	c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put(r.URL.String(), time.Now())
	})

	c.OnResponse(func(r *colly.Response) {
		url := r.Request.URL.String()
		start := r.Ctx.GetAny(url)
		if start != nil {
			duration := time.Now().Sub(start.(time.Time))
			if verbose {
				fmt.Printf("[%d] %5d ms : %s \n", r.StatusCode, duration.Milliseconds(), url)
			}
		}
	})

	q.AddURL(url)
	start := time.Now()
	q.Run(c)
	duration := time.Now().Sub(start)

	fmt.Printf("[URL] %5d ms : %s \n", duration.Milliseconds(), url)
}

type Args struct {
	Threads int
	Verbose bool
	Urls    []string
}

func usage() (Args, bool) {
	args := Args{0, false, []string{}}

	flag.IntVar(&args.Threads, "t", 1, "thread pool count")
	flag.BoolVar(&args.Verbose, "v", false, "verbose")
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
		Visit(url, args.Threads, args.Verbose)

		if args.Verbose {
			fmt.Println()
		}
	}
}
