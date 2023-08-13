package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {

	// Create a reverse proxy
	targetURL, _ := url.Parse("http://127.0.0.1:9091")
	targetURL2, _ := url.Parse("http://127.0.0.1:9090")

	counter := 0

	// Create an HTTP server to handle requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Modify the request before forwarding it

		tt := targetURL

		if counter%2 == 0 {
			tt = targetURL
		} else {
			tt = targetURL2

		}

		proxy := httputil.NewSingleHostReverseProxy(tt)
		proxy.ServeHTTP(w, r)

		counter++
	})

	fmt.Println("Starting server on port 8080")
	// Start the server
	http.ListenAndServe(":8080", nil)

}
