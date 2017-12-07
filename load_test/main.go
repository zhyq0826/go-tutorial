package main

import (
	"fmt"
	"net/http"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

func main() {
	rate := uint64(100) // per second
	duration := 4 * time.Second
	headers := make(http.Header)
  headers.Set("Origin", "http://wmem-test.wcar.net.cn")
  headers.Set("Sdate", "Wed, 06 Dec 2017 08:58:23 GMT")
  headers.Set("Host", "api_test_frontends")
  headers.Set("Accept-Language", "en-US,en;q=0.8")
  headers.Set("Accept-Encoding", "gzip, deflate")
  headers.Set("X-Forwarded-For", "124.207.11.6")
  headers.Set("X-Scheme","http")
  headers.Set("Content-Length","37")
  headers.Set("Accept", "*/*")
  headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.110 Safari/537.36")
  headers.Set("Date": "Wed, 06 Dec 2017 08:58:23 GMT")
  headers.Set("Connection", "close")
  headers.Set("Referer","http://wmem-test.wcar.net.cn/check")
  headers.Set("X-Real-Ip","124.207.11.6")
  headers.Set("Cookie","new_ford_pos=26dbb766b9e24f419fa91f13b967371c")
  headers.Set("Content-Type","application/x-www-form-urlencoded; charset=UTF-8")
  headers.Set("Authorization","Lb7VzSftregUAF2n:5626fe1276795e5fb17814ce1ae4266b")

	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    "http://localhost:11000/v3/base/home_latest_info",
    Body: []byte(),
    Header: headers,
	})
	attacker := vegeta.NewAttacker()
  attacker.

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration) {
		metrics.Add(res)
	}
	metrics.Close()

	fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}
