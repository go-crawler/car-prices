package fake

import (
	"math/rand"
)

var uas = [...]string{
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.63 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.111 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.1",
	"Mozilla/5.0 (Windows NT 6.2) AppleWebKit/536.3",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/535.24",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_0) AppleWebKit/536.3",
}

func GetUserAgent() string {
	n := rand.Intn(len(uas))
	return uas[n]
}