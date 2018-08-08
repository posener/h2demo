package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/posener/h2conn"
)

// A client for Go's HTTP2 echo server example at http2.golang.org/ECHO

func main() {
	// Create a client, that uses the HTTP PUT method.
	c := h2conn.Client{Method: http.MethodPut}

	// Connect to the HTTP2 server
	// The returned conn can be used to:
	//   1. Write - send data to the server.
	//   2. Read - receive data from the server.
	conn, resp, err := c.Connect(context.Background(), "https://http2.golang.org/ECHO")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Printf("Got: %d", resp.StatusCode)

	// Send time periodically to the server
	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Fprintf(conn, "It is now %v\n", time.Now())
		}
	}()

	// Read responses from the server to the stdout.
	_, err = io.Copy(os.Stdout, conn)
	if err != nil {
		log.Fatal(err)
	}
}
