

# go-sitecheck
- Site speed check  
- Content download speed corresponding to the tag[attribute] below

```
- link[href]
- script[src]
- img[src]
```

### Installing

```
$ git clone https://github.com/cdecl/go-sitecheck.git
$ cd go-sitecheck

$ go build
```

### Dependency
- https://github.com/gocolly/colly
- https://github.com/dustin/go-humanize

## Running the tests

- Usage

```
$ go-sitecheck.exe
Usage of go-sitecheck.exe:
  -     ver. 200428.0
  -H value
        add headers[]
  -json
        json output
  -m int
        request timeout sec (default 10)
  -t int
        thread pool count (default 1)
  -tag string
        add tag
  -v    verbose
  -x string
        request method (default "GET")
```

- Run

```
$ go-sitecheck http://httpbin.org/
[200]  2.1 MB [3043 ms] http://httpbin.org/

$ go-sitecheck -v http://httpbin.org/
[200]  9.6 kB [ 394 ms] http://httpbin.org/
[200]  1.8 kB [ 410 ms] https://fonts.googleapis.com/css?family=Open+Sans:400,700|Source+Code+Pro:300,600|Titillium+Web:400,600,700
[200]  154 kB [  22 ms] http://httpbin.org/flasgger_static/swagger-ui.css
[200]  1.4 MB [ 136 ms] http://httpbin.org/flasgger_static/swagger-ui-bundle.js
[200]  440 kB [  47 ms] http://httpbin.org/flasgger_static/swagger-ui-standalone-preset.js
[200]   86 kB [  15 ms] http://httpbin.org/flasgger_static/lib/jquery.min.js
[200]  2.1 MB [1113 ms] http://httpbin.org/

$ go-sitecheck -v -t 2 http://httpbin.org/
[200]  9.6 kB [ 388 ms] http://httpbin.org/
[200]  154 kB [  20 ms] http://httpbin.org/flasgger_static/swagger-ui.css
[200]  1.4 MB [ 135 ms] http://httpbin.org/flasgger_static/swagger-ui-bundle.js
[200]  440 kB [  52 ms] http://httpbin.org/flasgger_static/swagger-ui-standalone-preset.js
[200]   86 kB [  15 ms] http://httpbin.org/flasgger_static/lib/jquery.min.js
[200]  1.8 kB [ 381 ms] https://fonts.googleapis.com/css?family=Open+Sans:400,700|Source+Code+Pro:300,600|Titillium+Web:400,600,700
[200]  2.1 MB [ 854 ms] http://httpbin.org/

$ go-sitecheck -t 2 -json http://httpbin.org/
{"url":"http://httpbin.org/","status_code":200,"duration_ms":1942,"content_length":2120172,"timestamp":"2020-04-23T18:21:39+09:00"}

$ go-sitecheck -t 2 -tag httptest -json http://httpbin.org/
{"url":"http://httpbin.org/","status_code":200,"duration_ms":1095,"content_length":2120172,"tag":"httptest","timestamp":"2020-04-28T13:02:21+09:00"}

$ go-sitecheck -t 2 -json http://httpbin.org/ http://httpbin.org/get
{"url":"http://httpbin.org/","status_code":200,"duration_ms":886,"content_length":2120172,"timestamp":"2020-04-23T18:22:05+09:00"}
{"url":"http://httpbin.org/get","status_code":200,"duration_ms":667,"content_length":456,"timestamp":"2020-04-23T18:22:06+09:00"}

```


## License
This project is licensed under the MIT License 

