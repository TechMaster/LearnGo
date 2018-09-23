package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type Registry map[string][]string

func extractNameVersion(target *url.URL) (name, version string, err error) {
	path := target.Path
	// Trim the leading `/`
	if len(path) > 1 && path[0] == '/' {
		path = path[1:]
	}
	// Explode on `/` and make sure we have at least
	// 2 elements (service name and version)
	tmp := strings.Split(path, "/")
	if len(tmp) < 2 {
		return "", "", fmt.Errorf("Invalid path")
	}
	name, version = tmp[0], tmp[1]
	// Rewrite the request's path without the prefix.
	target.Path = "/" + strings.Join(tmp[2:], "/")
	return name, version, nil
}

// NewMultipleHostReverseProxy creates a reverse proxy that will randomly
// select a host from the passed `targets`
func NewMultipleHostReverseProxy(reg Registry) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		name, version, err := extractNameVersion(req.URL)
		if err != nil {
			log.Print(err)
			return
		}
		endpoints := reg[name+"/"+version]
		if len(endpoints) == 0 {
			log.Printf("Service/Version not found")
			return
		}
		req.URL.Scheme = "http"
		req.URL.Host = endpoints[rand.Int()%len(endpoints)]
	}
	return &httputil.ReverseProxy{
		Director: director,
	}
}

func main() {
	proxy := NewMultipleHostReverseProxy(Registry{
		"serviceone/v1": {"localhost:9091"},
		"serviceone/v2": {"localhost:9092"},
	})
	log.Fatal(http.ListenAndServe(":9090", proxy))
}