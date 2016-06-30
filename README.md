# User agent parsing
*useragent* is a library written in [golang](http://golang.org) to parse [user agent strings](http://useragentstring.com/).

# Usage
First install the library with:
```
go get xojoc.pw/useragent
```
*useragent* is simple to use. First parse a string with [useragent.Parse](http://godoc.org/xojoc.pw/useragent#Parse) and then access the fields of [useragent.UserAgent](http://godoc.org/xojoc.pw/useragent#UserAgent) for the required information. Example:
 * [Access fields](http://godoc.org/xojoc.pw/useragent#example-Parse--Access)

see [godoc](http://godoc.org/xojoc.pw/useragent) for the complete documentation.
# How it works?
          Lasciate ogne speranza, voi ch'intrate. -Dante
Parsing user agent strings is a hell. There is no standard for user agent strings, so *useragent* must use some heuristics. The site [http://www.useragentstring.com/](http://www.useragentstring.com/pages/useragentstring.php) has been invaluable during development. Some relevant links are also:

  * https://developer.mozilla.org/en-US/docs/Web/HTTP/Gecko_user_agent_string_reference
  * https://developer.chrome.com/multidevice/user-agent
  * https://developer.apple.com/library/safari/documentation/AppleApplications/Reference/SafariWebContent/OptimizingforSafarioniPhone/OptimizingforSafarioniPhone.html#//apple_ref/doc/uid/TP40006517-SW3
  * http://blogs.msdn.com/b/ieinternals/archive/2013/09/21/internet-explorer-11-user-agent-string-ua-string-sniffing-compatibility-with-gecko-webkit.aspx
  * http://docs.aws.amazon.com/silk/latest/developerguide/user-agent.html

for the supported user agents see:
  * browsers: [browser.go](browser.go)
  * crawlers: [crawler.go](crawler.go)

If you think *useragent* doesn't parse correctly a particular user agent string, just open an issue :).

# Example

see: example/main.go

```go
package main

import (
	"fmt"
	//"github.com/xojoc/useragent"
	"xojoc.pw/useragent" // 权威导入路径(import paths)
)

func main() {
	ua_list := [...]string{
		"Mozilla/5.0 (X11; Linux i686; rv:38.0) Gecko/20100101 Firefox/38.0",
		"Mozilla/5.0 (X11; U; Linux x86_64; sv-SE; rv:1.9.1.16) Gecko/20120714 IceCat/3.5.16 (like Firefox/3.5.16)",
		"Mozilla/5.0 (iPhone; U; CPU iPhone OS 5_1_1 like Mac OS X; en) AppleWebKit/534.46.0 (KHTML, like Gecko) CriOS/19.0.1084.60 Mobile/9B206 Safari/7534.48.3",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/50.0.2661.102 Chrome/50.0.2661.102 Safari/537.36",
	}

	for _, ua := range ua_list {
		p_ua := useragent.Parse(ua)
		println("=====")
		fmt.Print(p_ua)
	}
}
```

# Why this library?
*useragent* doesn't just split the user agent string and look for specific strings like other parsers, but it has specific parser for the most common browsers/crawlers and falls back to a generic parser for everything else. Its main features are:

 * Simple and stable API.
 * High precision in detection of the most common browsers/crawlers.
 * Detects mobile/tablet devices.
 * OS detection.
 * URL with more information about the user agent (usually its the home page).
 * [Security](http://godoc.org/xojoc.pw/useragent#Security) level detection when reported by browsers.


# Who?
*useragent* was written by Alexandru Cojocaru (http://xojoc.pw) and uses [blang/semver](https://github.com/blang/semver) to parse versions.

# License
*useragent* is released under the GPLv3 or later, see [COPYING](COPYING).
