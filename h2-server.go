package main

import (
	"log"
	"net/http"
)

func main() {
	// Create a server on port 8000
	// Exactly how you would run an HTTP1.1 server
	srv := &http.Server{Addr: ":8000", Handler: http.HandlerFunc(handle)}

	// Start the server with TLS, since we are running HTTP2 it must be run with TLS.
	// Exactly how you would run an HTTP1.1 server with TLS connection.
	log.Printf("Serving on https://0.0.0.0:8000")
	log.Fatal(srv.ListenAndServeTLS("server.crt", "server.key"))
}

func handle(w http.ResponseWriter, r *http.Request) {
	// Log the request protocol
	log.Printf("Got connection: %s", r.Proto)

	// Handle 2nd request, must be before push to prevent recursive calls.
	// Don't worry - Go protect us from recursive push by panicking.
	if r.URL.Path == "/2nd" {
		log.Println("Handling 2nd")
		w.Write([]byte("Hello Again!"))
		return
	}

	// Handle 1st request
	log.Println("Handling 1st")

	// Server push must be before response body is being written.
	// In order to check if the connection supports push, we should use
	// a type-assertion on the response writer.
	// If the connection does not support server push, or that the push fails we
	// just ignore it - server pushes are only here to improve the performance for HTTP2 clients.
	pusher, ok := w.(http.Pusher)
	if !ok {
		log.Println("Can't push to client")
	} else {
		err := pusher.Push("/2nd", nil)
		if err != nil {
			log.Printf("Failed push: %v", err)
		}
	}

	// Send response body
	w.Write([]byte("Hello"))
}
