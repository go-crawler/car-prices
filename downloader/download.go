package downloader

import (
	"log"
	"net/http"
	"io"

	"github.com/axgle/mahonia"

	"github.com/go-crawler/car-prices/fake"
)

func Get(url string) io.Reader {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("http.NewRequest err: %v", err)
	}

	req.Header.Add("User-Agent", fake.GetUserAgent())
	req.Header.Add("Referer", "https://car.autohome.com.cn")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("client.Do err: %v", err)
	}

	mah := mahonia.NewDecoder("gbk")
	return mah.NewReader(resp.Body)
}