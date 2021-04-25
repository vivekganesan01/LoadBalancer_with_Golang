package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	serverURLList = []string{
		"http://127.0.0.1:5000",
		"http://127.0.0.1:6001",
		"http://127.0.0.1:8000",
	}
	lastServedIndex = 0
)

func main() {
	http.HandleFunc("/", forwardRequest)
	log.Fatal(http.ListenAndServe(":9091", nil))
}

func forwardRequest(res http.ResponseWriter, req *http.Request) {
	url := getServerURL()
	rProxy := httputil.NewSingleHostReverseProxy(url)
	rProxy.ServeHTTP(res, req)

}

func getServerURL() *url.URL {
	mod := len(serverURLList)
	nextIndex := (lastServedIndex + 1) % mod
	url, _ := url.Parse(serverURLList[lastServedIndex])
	lastServedIndex = nextIndex
	return url
}
