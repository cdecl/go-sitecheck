

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
[200][ 2.1 MB][3330 ms] http://httpbin.org/

$ go-sitecheck -v http://httpbin.org/
[200][ 9.6 kB][ 493 ms] http://httpbin.org/
[200][ 1.8 kB][ 398 ms] https://fonts.googleapis.com/css?family=Open+Sans:400,700|Source+Code+Pro:300,600|Titillium+Web:400,600,700
[200][ 154 kB][  20 ms] http://httpbin.org/flasgger_static/swagger-ui.css
[200][ 1.4 MB][ 133 ms] http://httpbin.org/flasgger_static/swagger-ui-bundle.js
[200][ 440 kB][  48 ms] http://httpbin.org/flasgger_static/swagger-ui-standalone-preset.js
[200][  86 kB][  15 ms] http://httpbin.org/flasgger_static/lib/jquery.min.js
[200][ 2.1 MB][1158 ms] http://httpbin.org/

$ go-sitecheck -v -t 2 http://httpbin.org/
[200][ 9.6 kB][ 495 ms] http://httpbin.org/
[200][ 154 kB][  23 ms] http://httpbin.org/flasgger_static/swagger-ui.css
[200][ 1.4 MB][ 139 ms] http://httpbin.org/flasgger_static/swagger-ui-bundle.js
[200][ 440 kB][  49 ms] http://httpbin.org/flasgger_static/swagger-ui-standalone-preset.js
[200][  86 kB][  15 ms] http://httpbin.org/flasgger_static/lib/jquery.min.js
[200][ 1.8 kB][ 481 ms] https://fonts.googleapis.com/css?family=Open+Sans:400,700|Source+Code+Pro:300,600|Titillium+Web:400,600,700
[200][ 2.1 MB][1023 ms] http://httpbin.org/

$ go-sitecheck -t 2 -json http://httpbin.org/
{"url":"http://httpbin.org/","status_code":200,"duration_ms":1942,"content_length":2120172,"timestamp":"2020-04-23T18:21:39+09:00"}

$ go-sitecheck -t 2 -json http://httpbin.org/ http://httpbin.org/get
{"url":"http://httpbin.org/","status_code":200,"duration_ms":886,"content_length":2120172,"timestamp":"2020-04-23T18:22:05+09:00"}
{"url":"http://httpbin.org/get","status_code":200,"duration_ms":667,"content_length":456,"timestamp":"2020-04-23T18:22:06+09:00"}

```


## License
This project is licensed under the MIT License 

