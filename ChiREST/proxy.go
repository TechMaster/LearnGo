package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	//"runtime"
)

func main() {
	//runtime.GOMAXPROCS(2)
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "localhost:3333",
	})
	http.ListenAndServe(":8000", proxy)
}
