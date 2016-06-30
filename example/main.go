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
