

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

## Running the tests

- Usage

```
$ go-sitecheck
Usage of go-sitecheck:
  -json
        json output
  -t int
        thread pool count (default 1)
  -v    verbose
```

- Run

```
$ go-sitecheck http://httpbin.org/
[200]  1106 ms : http://httpbin.org/

$ go-sitecheck -v http://httpbin.org/
[200]   503 ms : http://httpbin.org/
[200]   413 ms : https://fonts.googleapis.com/css?family=Open+Sans:400,700|Source+Code+Pro:300,600|Titillium+Web:400,600,700
[200]    25 ms : http://httpbin.org/flasgger_static/swagger-ui.css
[200]     9 ms : http://httpbin.org/static/favicon.ico
[200]   141 ms : http://httpbin.org/flasgger_static/swagger-ui-bundle.js
[200]    59 ms : http://httpbin.org/flasgger_static/swagger-ui-standalone-preset.js
[200]    19 ms : http://httpbin.org/flasgger_static/lib/jquery.min.js
[200]  1223 ms : http://httpbin.org/

$ go-sitecheck -v -t 2 http://httpbin.org/
[200]   477 ms : http://httpbin.org/
[200]    25 ms : http://httpbin.org/flasgger_static/swagger-ui.css
[200]     8 ms : http://httpbin.org/static/favicon.ico
[200]   145 ms : http://httpbin.org/flasgger_static/swagger-ui-bundle.js
[200]    54 ms : http://httpbin.org/flasgger_static/swagger-ui-standalone-preset.js
[200]    20 ms : http://httpbin.org/flasgger_static/lib/jquery.min.js
[200]   411 ms : https://fonts.googleapis.com/css?family=Open+Sans:400,700|Source+Code+Pro:300,600|Titillium+Web:400,600,700
[200]   934 ms : http://httpbin.org/

$ go-sitecheck -t 2 -json http://httpbin.org/ http://httpbin.org/get
{"url":"http://httpbin.org/","status_code":200,"duration_ms":833,"timestamp":"2020-04-22T19:20:09+09:00"}

$ go-sitecheck -t 2 -json http://httpbin.org/ http://httpbin.org/get
{"url":"http://httpbin.org/","status_code":200,"duration_ms":3799,"timestamp":"2020-04-22T19:21:20+09:00"}
{"url":"http://httpbin.org/get","status_code":200,"duration_ms":625,"timestamp":"2020-04-22T19:21:24+09:00"}

```


## License

This project is licensed under the MIT License 

