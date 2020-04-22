

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
$ ./go-sitecheck
Usage of ./go-sitecheck:
  -t int
    	thread pool count (default 1)
  -v	verbose
```

- Run

```
$ ./go-sitecheck http://google.co.kr
[URL]   421 ms : http://google.co.kr 

$ ./go-sitecheck -v http://google.co.kr
[200]   103 ms : http://www.google.co.kr/logos/doodles/2020/earth-day-2020-6753651837108357.3-l.png 
[200]    38 ms : http://www.google.co.kr/textinputassistant/tia.png 
[URL]   538 ms : http://google.co.kr 

./go-sitecheck -t 2 -v http://google.co.kr
[200]    75 ms : http://www.google.co.kr/textinputassistant/tia.png 
[200]    98 ms : http://www.google.co.kr/logos/doodles/2020/earth-day-2020-6753651837108357.3-l.png 
[URL]   439 ms : http://google.co.kr 


$./go-sitecheck -t 4 -v  http://google.co.kr http://daum.net
[200]   129 ms : http://www.google.co.kr/textinputassistant/tia.png 
[200]   165 ms : http://www.google.co.kr/logos/doodles/2020/earth-day-2020-6753651837108357.3-l.png 
[URL]   505 ms : http://google.co.kr 

[200]    14 ms : https://search.daum.net/OpenSearch.xml 
[200]    36 ms : https://t1.daumcdn.net/adfit/static/ad.min.js 
[200]    42 ms : https://t1.daumcdn.net/adfit/static/ad-native.min.js?ver=201901201 
[200]     6 ms : https://t1.daumcdn.net/daumtop_deco/scripts/201905281455/daum_pctop.tm.min.js 
[200]    43 ms : https://t1.daumcdn.net/kas/static/ba.min.js 
[200]    29 ms : https://t1.daumcdn.net/daumtop_deco/scripts/202004211148/top.js 
[200]     3 ms : https://t1.daumcdn.net/b2/ssp/awsa.js?ver=20190920 
[200]     3 ms : https://t1.daumcdn.net/daumtop_chanel/op/20170315064553027.png 
[200]     3 ms : https://t1.daumcdn.net/daumtop_deco/banner/p_cafe_onlineclass.png 
[200]   900 ms : https://img2.daumcdn.net/thumb/C308x188/?fname=https://t1.daumcdn.net/section/oc/319485bacdd1424dbec1079e628d7115 
[200]   902 ms : https://img2.daumcdn.net/thumb/C308x188/?fname=https://t1.daumcdn.net/section/oc/c2b65349a1b04f5895c9e0f0c9716087 
[200]   905 ms : https://img2.daumcdn.net/thumb/C308x188/?fname=https://t1.daumcdn.net/section/oc/87fe9ca491e449ba926f5477e7007631 
[200]   904 ms : https://img2.daumcdn.net/thumb/C308x188/?fname=https://t1.daumcdn.net/section/oc/095a0389cfe4417b93b5189b07c613f4 
[200]     2 ms : https://t1.daumcdn.net/daumtop_deco/images/op/default.png 
[200]     4 ms : https://t1.daumcdn.net/daumtop_deco/images/op/default%402x.png 
[200]     7 ms : https://img2.daumcdn.net/thumb/C308x188/?fname=https://t1.daumcdn.net/daumtop_chanel/tromm/20200422113142732 
[200]     6 ms : https://img2.daumcdn.net/thumb/C308x188/?fname=https://t1.daumcdn.net/section/oc/d45008c262b24d30b81b2b63f7345d6c 
[200]     2 ms : https://t1.daumcdn.net/daumtop_deco/images/op/noimg_1x1.png 
[200]     2 ms : https://t1.daumcdn.net/daumtop_deco/images/top/2017/logo_foot.gif 
[URL]  1102 ms : http://daum.net 

```


## License

This project is licensed under the MIT License 

