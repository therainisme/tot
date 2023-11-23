package main

import (
	"flag"
	"net/http"
)

var port = flag.String("p", "8080", "the port to listen")

func main() {
	flag.Parse()

	http.HandleFunc("/logs", HandleProxyAvatarLogs)
	http.HandleFunc("/", HandleProxyAvatar)

	http.ListenAndServe(":"+*port, nil)
}