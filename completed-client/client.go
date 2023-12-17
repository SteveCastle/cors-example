package main

import "net/http"

// A web server that serves the static directory over HTTP on port 8081.

// This example demonstrates how to use the http package to create a web server




func main () {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8081", nil)
}