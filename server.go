package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloServerHttp(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server listening on http.\n"))
}

func HelloServerHttps(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server listening on https.\n"))
}

func main() {
	serverMuxHttp := http.NewServeMux()
	serverMuxHttp.HandleFunc("/", HelloServerHttp)

	serverMuxHttps := http.NewServeMux()
	serverMuxHttps.HandleFunc("/", HelloServerHttps)

	go func() {
		fmt.Println("Starting HTTP server")
		err := http.ListenAndServe(":8080", serverMuxHttp)
		if err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("Starting HTTPS server")
	err := http.ListenAndServeTLS(":8443", "server.crt", "server.key", serverMuxHttps)
	if err != nil {
		log.Fatal(err)
	}

}
