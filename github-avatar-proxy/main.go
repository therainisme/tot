package main

import (
	"flag"
	"log"
	"net/http"
)

var port = flag.String("p", "8080", "the port to listen")

func main() {
	flag.Parse()

	http.HandleFunc("/logs", HandleProxyAvatarLogs)
	http.HandleFunc("/", HandleProxyAvatar)

	log.Printf("listen on http://127.0.0.1:%s\n", *port)
	log.Printf("example: http://127.0.0.1:%s/therainisme.png\n", *port)

	http.ListenAndServe(":"+*port, nil)

	log.Printf("server exit\n")
}
