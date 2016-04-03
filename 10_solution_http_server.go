package main

import (
	"fmt"
	"net/http"
)

func startHTTPServer(port int) {
	http.ListenAndServe(fmt.Sprintf(":%d", port), http.FileServer(http.Dir("./static")))
}
